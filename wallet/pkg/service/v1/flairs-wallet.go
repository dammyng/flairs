package v1

import (
	"context"
	"time"
	v1internals "wallet/internals/v1"
	v1 "wallet/pkg/api/v1"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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

func (f *flairsWalletServer) AddNewWallet(ctx context.Context, req *v1.NewWalletRequest) (*v1.AddWalletResponse, error) {

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

	if req.UserId != claims.UserID {
		return nil, status.Error(codes.Unauthenticated, "Error fetching user record ")
	}

	ID := uuid.NewV4().String()

	newWallet := v1.Wallet{
		AccountBal:    req.AccountBal,
		LedgerBal:     req.LedgerBal,
		Currency:      req.Currency,
		ID:            ID,
		WalletType:    req.WalletType,
		UserId:        req.UserId,
		Memo:          req.Memo,
		Name:          req.Name,
		Status:        req.Status,
		DateCreated:   time.Now().Format(time.RFC3339),
		DateBalUpdate: time.Now().Format(time.RFC3339),
		LastUpdate:    time.Now().Format(time.RFC3339),
	}

	err = f.Db.CreateWallet(&newWallet)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create new wallet-> "+err.Error())
	}

	// Response
	return &v1.AddWalletResponse{
		ID: ID,
	}, nil
}

func (f *flairsWalletServer) AddWalletTerm(ctx context.Context, req *v1.NewWalletTermRequest) (*v1.AddWalletTermResponse, error) {
	return nil, nil
}

func (f *flairsWalletServer) AddWalletType(ctx context.Context, req *v1.NewWalletTypeRequest) (*v1.AddWalletTypeResponse, error) {
	return nil, nil
}

func (f *flairsWalletServer) GetOneWallet(ctx context.Context, req *v1.GetOneWalletReq) (*v1.GetOneWalletRes, error) {
	w, err := f.Db.GetWallet(&v1.GetOneWalletReq{WalletId: req.WalletId})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Wallet not found")
	}
	return &v1.GetOneWalletRes{
		AccountBal: w.AccountBal,
		Currency: w.Currency,
		DateBalUpdate: w.DateBalUpdate,
		DateCreated: w.DateCreated,
		LastUpdate: w.LastUpdate,
		LedgerBal: w.LedgerBal,
		Memo: w.Memo,
		Name: w.Name,
		Status: w.Status,
		WalletType: w.WalletType,
	}, nil
}

func (f *flairsWalletServer) GetMyWallets(ctx context.Context, req *v1.GetMyWalletsRequest) (*v1.WalletsResponse, error) {

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

	if req.UserId != claims.UserID {
		return nil, status.Error(codes.Unauthenticated, "Error fetching user record ")
	}

	_ws, err := f.Db.GetUserWallets(req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Get my wallets-> "+err.Error())
	}

	var ws v1.WalletsResponse

	for _, v := range _ws {
		ws.Wallets = append(ws.Wallets, &v)
	}
	// Response
	return &ws, nil
}

func (f *flairsWalletServer) Transact(ctx context.Context, req *v1.PerformTransactionReq) (*v1.PerformTransactionRes, error) {

	// Get wallet
	w, err := f.Db.GetWallet(&v1.GetOneWalletReq{WalletId: req.WalletID})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Wallet not found")
	}
	// update the balance
	newBal := w.AccountBal + req.Amount
	// gorm save
	err = f.Db.UpdateWallet(&w, &v1.Wallet{AccountBal: newBal})
	// return wallet ID,
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unable to complete")
	}
	return &v1.PerformTransactionRes{
		ID: req.WalletID,
	}, nil
}
