package v1

import (
	v1 "auth/pkg/api/v1"
	"context"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddNewUser initializes a new user with email address
func (f *flairsServiceServer) AddNewUser(ctx context.Context, req *v1.AddNewUserRequest) (*v1.AddNewUserResponse, error) {
	ID := uuid.NewV4()
	req.ID = ID.String()
	
	err := f.Session.Create(v1.User{ID: req.ID, Email: req.Email}).Error
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Users-> "+err.Error())
	}
	return &v1.AddNewUserResponse{
		Api: apiVersion,
		ID:  req.ID,
	}, nil
}
