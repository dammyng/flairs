package v1

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	v1internals "wallet/internals/v1"
	v1 "wallet/pkg/api/v1"
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

func (f *flairsWalletServer) AddNewWallet(ctx context.Context, req *v1.NewWalletRequest) (*v1.AddWalletResponse, error) {
	ID := uuid.NewV4().String()

	newWallet := v1.Wallet{
		AccountBal:    0.00,
		LedgerBal:     0.00,
		Currency:      req.Currency,
		ID:            ID,
		WalletType:    req.WalletType,
		UserId:        req.UserId,
		Memo:          req.Memo,
		Name:          req.Name,
		TermID:        "1",
		Status:        true,
		DateCreated:   time.Now().Format(time.RFC3339),
		DateBalUpdate: time.Now().Format(time.RFC3339),
		LastUpdate:    time.Now().Format(time.RFC3339),
	}

	err := f.Db.CreateWallet(&newWallet)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create new wallet-> "+err.Error())
	}

	// Response
	return &v1.AddWalletResponse{
		ID: ID,
	}, nil
}
