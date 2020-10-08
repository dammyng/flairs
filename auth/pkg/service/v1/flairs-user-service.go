package v1

import (
	v1 "auth/pkg/api/v1"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddNewUser initializes a new user with email address
func (f *flairsServiceServer) AddNewUser(ctx context.Context, req *v1.AddNewUserRequest) (*v1.AddNewUserResponse, error) {
	err := f.Session.Create(req).Error
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Users-> "+err.Error())
	}
	return &v1.AddNewUserResponse{
		Api: apiVersion,
		ID:  req.ID,
	}, nil
}
