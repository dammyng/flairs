package v1

import (
	v1 "auth/pkg/api/v1"
	v1helper "auth/pkg/helper/v1"
	"context"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"shared/events"

	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AddNewUser initializes a new user with email address
func (f *flairsServiceServer) AddNewUser(ctx context.Context, req *v1.AddNewUserRequest) (*v1.AddNewUserResponse, error) {
	if !v1helper.IsEmailValid(req.Email) {
		return nil, status.Error(codes.InvalidArgument, "Invalid entry")
	}

	user := v1helper.DecodeToSQLUser(req)

	// Check if user Already exist
	if _, err := f.Db.FindUser(user); err != gorm.ErrRecordNotFound {
		return nil, status.Error(codes.AlreadyExists, "A Flairs account with this email already exists. Please try logging in.")
	}

	// User ID
	ID := uuid.NewV4().String()
	user.ID = ID

	// User Referrer
	ref := req.Ref
	if len(ref) > 0 {
		user.Referrer = ref
	}

	// Create user referrer code
	refcode := v1helper.RandString(15)
	user.RefCode = refcode

	// Database call
	err := f.Db.CreateUser(user)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to insert into Users-> "+err.Error())
	}

	// Token - Store token in redis cache
	token := v1helper.RandInt(6)
	log.Println(token)
	_, err = f.RedisConn.Do("HMSET", "email:verification", user.Email, token)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.RedisConn.Do("HMSET", "password:reset", user.Email, token)
	if err != nil {
		log.Fatal(err)
	}

	// Response
	return &v1.AddNewUserResponse{
		ID: ID,
	}, nil
}

func (f *flairsServiceServer) LoginUser(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {

	// Check if user Already exist
	user, err := f.Db.FindUser(&v1.User{Email: req.Email})
	if user == nil {
		return nil, status.Error(codes.NotFound, "A Flairs account with this email does not exist exists.")
	}

	if user == nil && err != nil {
		return nil, status.Error(codes.NotFound, "A Flairs account with this email does not exist exists.")
	}
	emailVerifiedAt := user.EmailVerifiedAt

	if emailVerifiedAt.IsZero() {
		return nil, status.Error(codes.Unauthenticated, "Unverified email address")
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.Internal, "Unverified email address")
	}

	expirationTime := time.Now().Add(24 * 60 * time.Minute)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secrek_key"))
	if err != nil {
		return nil, status.Error(codes.Internal, "Aist exists.")
	}

	res := &v1.LoginResponse{
		Token: tokenString,
		User: &v1.Profile{
			ID: user.ID,
		},
	}

	return res, nil
}

func (f *flairsServiceServer) ReadUserBy(ctx context.Context, req *v1.ReadUserByRequest) (*v1.ReadUserByResponse, error) {
	return nil, nil
}

func (f *flairsServiceServer) ResetUserPassword(ctx context.Context, req *v1.ResetPasswordRequest) (*v1.CustomResponse, error) {
	u := v1helper.DecodeToSQLUser(req)
	user, err := f.Db.FindUser(u)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
	}

	emailVerfiedAt := user.EmailVerifiedAt
	if emailVerfiedAt.IsZero() {
		return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
	}

	token := v1helper.RandInt(6)
	_, err = f.RedisConn.Do("HMSET", "password:reset", user.Email, token)

	//
	//Rabbit send the mail

	// TODO how to handle this error
	if err != nil {
		return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
	}
	return &v1.CustomResponse{
		Message: "Successful",
		Request: "Reset password",
	}, nil
}

func (f *flairsServiceServer) SetUserPassword(ctx context.Context, req *v1.SetPasswordRequest) (*v1.CustomResponse, error) {
	u := v1helper.DecodeToSQLUser(req)

	md, ok := metadata.FromIncomingContext(ctx)
	log.Println(ok)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	}
	authorization := md.Get("Authorization")[0]
	if authorization == "" {
		return nil, status.Error(codes.Unauthenticated, "Invalid authorization token ")
	}

	token, err := redis.String(f.RedisConn.Do("HGET", "password:reset", req.Email))
	if err != nil {
		return nil, status.Error(codes.Internal, "Invalid token string "+err.Error())
	}
	if token == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid token string "+err.Error())
	}

	if match := token == authorization; match {
		user, err := f.Db.FindUser(u)
		u.ID = user.ID
		if err != nil {
			return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, status.Error(codes.Internal, "Service failed"+err.Error())
		}
		err = f.Db.UpdateUser(u, &v1.User{Password: hashedPass})
		if err != nil {
			return nil, status.Error(codes.Internal, "Service failed"+err.Error())
		}
		redis.Int(f.RedisConn.Do("HDEL", "password:reset", req.Email))
		res := &v1.CustomResponse{
			Message: "Successfully Add Password",
			Request: "add_password",
		}

		return res, nil
	}
	return nil, status.Error(codes.InvalidArgument, "Wrong token string")

}

func (f *flairsServiceServer) UpdateUserProfile(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {

	u := v1helper.DecodeToSQLUser(req.Profile)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to get metadata")
	}
	authorization := md.Get("Authorization")[0]
	if authorization == "" {
		return nil, status.Error(codes.Unauthenticated, "Invalid authorization token ")
	}

	claims := &Claims{}
	err := DecodeJwt(authorization, claims)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid token")

		}
		return nil, status.Errorf(codes.Unauthenticated, "Token inaccessible")
	}
	_, err = f.Db.FindUser(&v1.User{ID: req.Id})

	if err != nil {
		return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
	}

	if req.Id != claims.UserID {
		return nil, status.Error(codes.Unauthenticated, "Error fetching user record ")
	}

	uu := &v1.User{ID: req.Id}
	err = f.Db.UpdateUser(uu, u)

	profileComplete := IsProfileCompled(uu)
	err = f.Db.UpdateUser(uu, &v1.User{IsProfileCompleted: profileComplete})

	if err != nil {
		return nil, status.Error(codes.Internal, "Service failed"+err.Error())
	}
	res := &v1.UpdateUserResponse{
		Id: req.Id,
	}

	return res, nil
}

func (f *flairsServiceServer) ValidateUserEmail(ctx context.Context, req *v1.ValidateEmailRequest) (*v1.CustomResponse, error) {

	u := v1helper.DecodeToSQLUser(req)

	token, err := redis.String(f.RedisConn.Do("HGET", "email:verification", req.Email))
	if err != nil {
		return nil, status.Error(codes.Internal, "Invalid token string "+err.Error())
	}
	if token == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid token string "+err.Error())
	}

	if match := token == req.Token; match {
		user, err := f.Db.FindUser(u)
		u.ID = user.ID
		if err != nil {
			return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
		}

		err = f.Db.UpdateUser(u, &v1.User{EmailVerifiedAt: time.Now().Format(time.RFC3339)})

		if err != nil {
			return nil, status.Error(codes.Internal, "Invalid token string "+err.Error())
		}
		redis.Int(f.RedisConn.Do("HDEL", "email:verification", req.Email))

		// Token - Send token to email with Rabbit
		msg := events.CreateDefWallet{
			URL:    "http://localhost:9000/v1/wallet",
			UserID: user.ID,
		}
		f.EventEmitter.Emit(&msg, "auth")
	} else {
		return nil, status.Error(codes.InvalidArgument, "Wrong token string")
	}

	return &v1.CustomResponse{
		Message: "Successfully verified email",
		Request: "verify_email",
	}, nil
}
