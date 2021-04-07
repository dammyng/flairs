package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"


	amqp "shared/events/amqp"
	v1internals "transaction/internals/v1"
	v1 "transaction/pkg/api/v1"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
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

func (f *flairsTransactionServer) AddNewTransaction(ctx context.Context, req *v1.NewTransactionReq) (*v1.NewTransactionRes, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, NoAuthMetaDataError
	}
	authorization := md.Get("Authorization")[0]
	if authorization == "" {
		return nil, InvalidTokenError
	}

	claims := &Claims{}

	err := DecodeJwt(authorization, claims)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err

		}
		return nil, WrongTokenStruct
	}

	newTransaction := v1.Transaction{
		ID:         uuid.NewV4().String(),
		Amount:     req.Amount,
		Memo:       req.Memo,
		WalletId:   req.WalletId,
		Status:     req.Status,
		Currency:   req.Currency,
		TxRef:      req.TxRef,
		TransType:  v1.TransactionType(req.TransType),
		CustomerId: req.CustomerId,
		Source:     req.Source,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}

	id, err := f.Db.CreateTransaction(&newTransaction)
	if err != nil {
		return nil, InternalError
	}

	return &v1.NewTransactionRes{Id: id}, err

}

func (f *flairsTransactionServer) GetTransaction(ctx context.Context, req *v1.GetTransactionRequest) (*v1.TransactionResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, NoAuthMetaDataError
	}
	authorization := md.Get("Authorization")[0]
	if authorization == "" {
		return nil, InvalidTokenError
	}

	claims := &Claims{}

	err := DecodeJwt(authorization, claims)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err

		}
		return nil, WrongTokenStruct
	}

	transaction, err := f.Db.GetTransaction(&v1.Transaction{ID: req.Id})

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, WalletNotFoundError
	}

	if err != nil {
		return nil, InternalError
	}

	return &v1.TransactionResponse{
		Transaction: transaction,
	}, err
}

func (f *flairsTransactionServer) UpdateTransaction(ctx context.Context, req *v1.UpdateTransactionsRequest) (*v1.UpdateTransactionResponse, error) {
	return nil, nil
}

func (f *flairsTransactionServer) GetMyTransactions(ctx context.Context, req *v1.GetMyTransactionsRequest) (*v1.TransactionsResponse, error) {
	return nil, nil
}

func (f *flairsTransactionServer) GetWalletTransactions(ctx context.Context, req *v1.GetWalletTransactionsRequest) (*v1.WalletBalanceResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, NoAuthMetaDataError
	}
	authorization := md.Get("Authorization")[0]
	if authorization == "" {
		return nil, InvalidTokenError
	}

	claims := &Claims{}

	err := DecodeJwt(authorization, claims)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err

		}
		return nil, WrongTokenStruct
	}

	transactions, err := f.Db.GetWalletTransactions(&v1.Transaction{WalletId: req.WalletId})

	if err != nil {
		return nil, InternalError
	}

	var wb v1.WalletBalanceResponse

	for _, v := range transactions {
		wb.Transactions = append(wb.Transactions, &v)
	}
	wb.Balance = evaluateBalance(wb.Transactions)

	return &wb, nil
}
func (f *flairsTransactionServer) TransactFtDeposite(ctx context.Context, req *v1.TfDepositeRequest) (*v1.TransactResponse, error) {

	reqURL, _ := url.Parse(fmt.Sprintf("https://api.flutterwave.com/v3/transactions/%v/verify", req.TxRef))
	flutterReq := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json; charset=UTF-8"},
			"Authorization": {"Bearer " + os.Getenv("FlutterSecret")},
		},
	}
	result, err := HttpReq(flutterReq)

	if err != nil {
		return nil, err
	}
	if result.StatusCode > 299 {
		return nil, err
	}
	var response map[string]string
	var data map[string]string

	err = json.NewDecoder(result.Body).Decode(response)
	err = json.NewDecoder(strings.NewReader(response["data"])).Decode(data)

	amt, _ := strconv.ParseFloat(data["Amount"], 8)

	newTransaction := v1.Transaction{
		ID:         uuid.NewV4().String(),
		Amount:     amt,
		Memo:       "Wallet funds",
		WalletId:   req.WalletId,
		Status:     response["status"],
		Currency:   "",
		TxRef:      req.TxRef,
		TransType:  v1.Transaction_CREDIT,
		CustomerId: "",
		Source:     "FW",
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}

	_, err = f.Db.CreateTransaction(&newTransaction)
	if err != nil {
		return nil, InternalError
	}

	return &v1.TransactResponse{Amount: float32(amt)}, err
}

func (f *flairsTransactionServer) TransactWalletTransfer(ctx context.Context, req *v1.TransferRequest) (*v1.TransactResponse, error) {
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

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var netClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	res, err := netClient.Do(req)

	// check for response error
	if err != nil {
		log.Fatal("Error:", err)
		return nil, err
	}
	return res, err
}

func evaluateBalance(ts []*v1.Transaction) float64 {
	total := 0.0
	for _, v := range ts {
		if v.TransType == v1.Transaction_CREDIT && v.Status != "failed" {
			total += v.Amount
		}
		if v.TransType == v1.Transaction_DEBIT && v.Status != "failed" {
			total += v.Amount
		}
	}
	return total
}
