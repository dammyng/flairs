package v1

import (
	v1 "auth/pkg/api/v1"
	v1helper "auth/pkg/helper/v1"
	"context"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AddNewUser initializes a new user with email address
func (f *flairsServiceServer) AddNewUser(ctx context.Context, req *v1.AddNewUserRequest) (*v1.AddNewUserResponse, error) {
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

	// Token - Send token to email with Rabbit

	// Response
	return &v1.AddNewUserResponse{
		Api: apiVersion,
		ID:  ID,
	}, nil
}

func (f *flairsServiceServer) LoginUser(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {

	return nil, nil
}

func (f *flairsServiceServer) ReadUserBy(ctx context.Context, req *v1.ReadUserByRequest) (*v1.ReadUserByResponse, error) {
	return nil, nil
}

func (f *flairsServiceServer) ResetUserPassword(ctx context.Context, req *v1.ResetPasswordRequest) (*v1.CustomResponse, error) {
	return nil, nil
}

func (f *flairsServiceServer) SetUserPassword(ctx context.Context, req *v1.SetPasswordRequest) (*v1.CustomResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	log.Println(md.Get("Authorization"))
	return nil, nil
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
		redis.Int(f.RedisConn.Do("HDEL", "password:reset", req.Email))

	} else {
		return nil, status.Error(codes.InvalidArgument, "Wrong token string")
	}
	return &v1.CustomResponse{
		Api: "v1",
		Message: "Successfully verified email",
		Request: "verify_email",
	}, nil
}

