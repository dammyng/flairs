package v1

import (
	v1internals "wallet/internals/v1"
	v1 "wallet/pkg/api/v1"

	"context"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type flairsWalletServer struct {
	Db v1internals.DatabaseHandler
}

// Claims jwt custom authentication claims
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// NewflairsWalletServer creates ToDo service
func NewflairsWalletServer(db v1internals.DatabaseHandler) v1.WalletServiceServer {
	return &flairsWalletServer{Db: db}
}

// DecodeJwt - decodes JWT token from request
func DecodeJwt(token string, claims *Claims) error {
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secrek_key"), nil
	})
	return err
}

func (f *flairsWalletServer) AddNewWallet(ctx context.Context, req *v1.NewWalletRequest)( *v1.AddWalletResponse, error) {
	return nil, status.Error(codes.Internal, "")
}
