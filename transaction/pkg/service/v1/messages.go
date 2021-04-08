package v1

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


var(
	NoAuthMetaDataError = status.Errorf(codes.DataLoss, "failed to get authentication token - 50")
	InvalidTokenError = status.Error(codes.Unauthenticated, "Invalid authorization token - 40")
	WrongTokenStruct = status.Errorf(codes.Unauthenticated, "Failed to handle token")
	UserIDClaimIDError = status.Error(codes.Unauthenticated, "Error fetching user record ")
	WalletNotFoundError = status.Errorf(codes.NotFound, "Wallet not found")
	InternalError = status.Errorf(codes.Internal, "Internal service error")
)

func flutterError(data interface{}) error  {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("%v", data ) )
}