package v1

import (
	"context"

	//"io/ioutil"
	"log"
	"net/http"
	
	amqp "shared/events/amqp"
	v1internals "transaction/internals/v1"
	v1 "transaction/pkg/api/v1"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type flairsTransactionServer struct {
	Db           v1internals.DatabaseHandler
	EventEmitter amqp.EventEmitter                                                              
}

// Claims jwt custom authentication claims
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// NewflairsTransactionServer creates ToDo service
func NewflairsTransactionServer(db v1internals.DatabaseHandler, eventEmitter amqp.EventEmitter) v1.FlairsTransactionServiceServer {
	return &flairsTransactionServer{Db: db,
		EventEmitter: eventEmitter,
	}
}

// DecodeJwt - decodes JWT token from request
func DecodeJwt(token string, claims *Claims) error {
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secrek_key"), nil
	})
	return err
}

func (f *flairsTransactionServer) AddNewTransaction(ctx context.Context, req *v1.Transaction) (*v1.NewTransactionRes, error) {

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

	
	
	_ = uuid.NewV4().String()

	return nil, status.Error(codes.InvalidArgument, "Invalid transaction type")
}


func (f *flairsTransactionServer) GetTransaction(ctx context.Context, req *v1.GetTransactionRequest) (*v1.TransactionResponse, error) {
	return nil, nil
}

func (f *flairsTransactionServer) UpdateTransaction(ctx context.Context, req *v1.UpdateTransactionsRequest) (*v1.UpdateTransactionResponse, error) {
	return nil, nil
}

func (f *flairsTransactionServer) GetMyTransactions(ctx context.Context, req *v1.GetMyTransactionsRequest) (*v1.TransactionsResponse, error) {
	return nil, nil
}

func (f *flairsTransactionServer) GetWalletTransactions(ctx context.Context, req *v1.GetWalletTransactionsRequest) (*v1.WalletBalanceResponse, error) {
	return nil, nil
}





//{status: "successful", customer: {…}, transaction_id: 1695241, tx_ref: "hooli-tx-1920bbtyt", flw_ref: "FLW-MOCK-6439899760b3449a2db802decd80594f", …}
//amount: 100
//currency: "NGN"
//customer: {name: "yemi desola", email: "user@gmail.com", phone_number: "08102909304"}
//flw_ref: "FLW-MOCK-6439899760b3449a2db802decd80594f"
//status: "successful"
//transaction_id: 1695241
//tx_ref: "hooli-tx-1920bbtyt"
//curl --location --request GET 'https://api.flutterwave.com/v3/transactions/123456/verify' \
//--header 'Content-Type: application/json' \
//--header 'Authorization: Bearer {{SEC_KEY}}'



func HttpReq(req *http.Request) (*http.Response, error) {

	// send an HTTP request using `req` object
	res, err := http.DefaultClient.Do(req)

	// check for response error
	if err != nil {
		log.Fatal("Error:", err)
		return nil, err
	}
	return res, err
}
