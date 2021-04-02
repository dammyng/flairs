package v1

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


var(
	NoAuthMetaDataError = status.Errorf(codes.DataLoss, "failed to get authentication token - 50")
	InvalidTokenError = status.Error(codes.Unauthenticated, "Invalid authorization token - 40")
	WrongTokenStruct = status.Errorf(codes.Unauthenticated, "Failed to handle token")
)