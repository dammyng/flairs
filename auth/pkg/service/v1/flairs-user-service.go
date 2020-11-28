package v1

import (
	v1 "auth/pkg/api/v1"
	v1helper "auth/pkg/helper/v1"
	"context"
	"encoding/hex"
	"log"
	"os"
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
		return nil, status.Error(codes.InvalidArgument, "Invalid entry - Enter a valid email")
	}

	user := v1helper.DecodeToSQLUser(req)

	// Check if user Already exist
	if _, err := f.Db.FindUser(user); err != gorm.ErrRecordNotFound {
		return nil, status.Error(codes.AlreadyExists, "A Flairs account with this email already exists. Please try logging in.")
	}

	// User ID
	ID := uuid.NewV4()
	user.ID = ID.String()

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
		return nil, status.Error(codes.Internal, "Something went wrong"+err.Error())
	}

	// Token - Store token in redis cache
	token := v1helper.RandInt(6)
	_, err = f.RedisConn.Do("HMSET", "email:verification", user.Email, token)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.RedisConn.Do("HMSET", "password:reset", user.Email, token)
	if err != nil {
		log.Fatal(err)
	}

	msg := events.UserCreatedEvent{
		ID:    hex.EncodeToString(ID.Bytes()),
		Email: user.Email,
		Token: token,
	}
	f.EventEmitter.Emit(&msg, "auth")

	// Response
	return &v1.AddNewUserResponse{
		ID: ID.String(),
	}, nil
}

func (f *flairsServiceServer) LoginUser(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {

	// Check if user Already exist
	user, err := f.Db.FindUser(&v1.User{Email: req.Email})
	if user == nil {
		return nil, status.Error(codes.NotFound, "A Flairs account with this email does not exist exists.")
	}

	if user == nil && err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong verifying user.")
	}
	emailVerifiedAt := user.EmailVerifiedAt

	if emailVerifiedAt.IsZero() {
		return nil, status.Error(codes.Unauthenticated, "Unverified email address")
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong verifying user.")
	}

	expirationTime := time.Now().Add(24 * 60 * time.Minute)

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_FLAIRS_KEY")))
	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong verifying user")
	}

	res := &v1.LoginResponse{
		Token: tokenString,
		User: &v1.Profile{
			ID:    user.ID,
			Email: user.Email,
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

	emailVerifiedAt := user.EmailVerifiedAt
	if emailVerifiedAt.IsZero() {
		return nil, status.Error(codes.NotFound, "You cannot reset password for unverified user account"+err.Error())
	}

	token := v1helper.RandInt(6)
	_, err = f.RedisConn.Do("HMSET", "password:reset", user.Email, token)

	//Rabbit send the mail
	id := uuid.NewV4()
	//fmt.Println("sending reset event")
	msg := events.PasswordReset{
		ID:    hex.EncodeToString(id.Bytes()),
		Email: user.Email,
		Token: token,
	}

	err = f.EventEmitter.Emit(&msg, "auth")

	// TODO how to handle this error
	if err != nil {
		return nil, status.Error(codes.NotFound, "Could not send mail to user"+err.Error())
	}
	return &v1.CustomResponse{
		Message: "Successful",
		Request: "Request_Password_Reset",
	}, nil
}

func (f *flairsServiceServer) SendValidationMail(ctx context.Context, req *v1.SendValidationMailRequest) (*v1.CustomResponse, error) {
	u := v1helper.DecodeToSQLUser(req)
	user, err := f.Db.FindUser(u)
	if err != nil {
		return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
	}
	emailVerifiedAt := user.EmailVerifiedAt
	if !emailVerifiedAt.IsZero() {
		return nil, status.Error(codes.NotFound, "You cannot validate an already validated email "+err.Error())
	}

	token := v1helper.RandInt(6)
	_, err = f.RedisConn.Do("HMSET", "email:verification", user.Email, token)

	msg := events.UserCreatedEvent{
		Email: user.Email,
		Token: token,
	}
	err = f.EventEmitter.Emit(&msg, "auth")

	// TODO how to handle this error
	if err != nil {
		return nil, status.Error(codes.NotFound, "Could not send mail to user"+err.Error())
	}
	return &v1.CustomResponse{
		Message: "Successful",
		Request: "Request_Email_Validation",
	}, nil
}

func (f *flairsServiceServer) SetUserPassword(ctx context.Context, req *v1.SetPasswordRequest) (*v1.CustomResponse, error) {
	u := v1helper.DecodeToSQLUser(req)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to retrive authentication token")
	}
	authorization := md.Get("Authorization")[0]
	if authorization == "" {
		return nil, status.Error(codes.Unauthenticated, "Invalid authorization token")
	}

	token, err := redis.String(f.RedisConn.Do("HGET", "password:reset", req.Email))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to retrive authentication token "+err.Error())
	}
	if token == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty authentication token "+err.Error())
	}

	if match := token == authorization; match {
		user, err := f.Db.FindUser(u)
		u.ID = user.ID
		if err != nil {
			return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, status.Error(codes.Internal, "Something went wrong"+err.Error())
		}
		err = f.Db.UpdateUser(u, &v1.User{Password: hashedPass})
		if err != nil {
			return nil, status.Error(codes.Internal, "Something went wrong"+err.Error())
		}
		redis.Int(f.RedisConn.Do("HDEL", "password:reset", req.Email))
		res := &v1.CustomResponse{
			Message: "Successfully",
			Request: "set_password",
		}

		return res, nil
	}
	return nil, status.Error(codes.InvalidArgument, "Invalid authorization token")
}

func (f *flairsServiceServer) UpdateUserProfile(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserResponse, error) {
	u := v1helper.DecodeToSQLUser(req.Profile)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "failed to retrive authentication token")
	}
	authorization := md.Get("Authorization")[0]
	if authorization == "" {
		return nil, status.Error(codes.Unauthenticated, "Invalid authorization token ")
	}

	claims := &Claims{}
	err := DecodeJwt(authorization, claims)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, status.Errorf(codes.Unauthenticated, "Could not authenticate this request")

		}
		return nil, status.Errorf(codes.Unauthenticated, "Could not authenticate - Token inaccessible")
	}
	_, err = f.Db.FindUser(&v1.User{ID: req.Id})

	if err != nil {
		return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
	}

	if req.Id != claims.UserID {
		return nil, status.Error(codes.Unauthenticated, "Invalid authentication token")
	}

	uu := &v1.User{ID: req.Id}
	err = f.Db.UpdateUser(uu, u)

	profileComplete := IsProfileCompled(uu)
	err = f.Db.UpdateUser(uu, &v1.User{IsProfileCompleted: profileComplete})

	if err != nil {
		return nil, status.Error(codes.Internal, "Something went wrong "+err.Error())
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
		return nil, status.Error(codes.Internal, "failed to retrive authentication token "+err.Error())
	}
	if token == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid authentication token "+err.Error())
	}

	if match := token == req.Token; match {
		user, err := f.Db.FindUser(u)
		u.ID = user.ID
		if err != nil {
			return nil, status.Error(codes.NotFound, "Error fetching user record "+err.Error())
		}

		err = f.Db.UpdateUser(u, &v1.User{EmailVerifiedAt: time.Now().Format(time.RFC3339)})

		if err != nil {
			return nil, status.Error(codes.Internal, "Something went wrong "+err.Error())
		}
		redis.Int(f.RedisConn.Do("HDEL", "email:verification", req.Email))



		// Token required to create default wallet once email is a validated
		expirationTime := time.Now().Add(24 * 60 * time.Minute)

		claims := &Claims{
			UserID: user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_FLAIRS_KEY")))

		msg := events.CreateDefWallet{
			UserID: user.ID,
			Token:  tokenString,
		}
		f.EventEmitter.Emit(&msg, "auth")
	} else {
		return nil, status.Error(codes.InvalidArgument, "Invalid authentication token")
	}

	return &v1.CustomResponse{
		Message: "Successfully",
		Request: "verify_email",
	}, nil
}
