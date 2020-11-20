package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	//"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"shared/events"
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

func (f *flairsTransactionServer) AddnewTransaction(ctx context.Context, req *v1.NewTransactionReq) (*v1.NewTransactionRes, error) {

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

	var newT = v1.Transaction{
		ID:         ID,
		CustomerId: req.UserId,
		Memo:       req.UserMemo,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
		Currency:   "",
	}

	switch req.TransactionType {
	case 0:
		reqURL, _ := url.Parse(fmt.Sprintf("https://api.flutterwave.com/v3/transactions/%v/verify", req.ThirdPartyID))
		flutterReq := &http.Request{
			Method: "GET",
			URL:    reqURL,
			Header: map[string][]string{
				"Content-Type":  {"application/json; charset=UTF-8"},
				"Authorization": {"Bearer " + os.Getenv("FlutterSecret")},
			},
			//Body: reqBody,
		}
		res, err := HttpReq(flutterReq)
		if err != nil {
			return nil, err
		}
		if res.StatusCode > 299 {
			return nil, err
		}
		var result map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)
		// close response body
		res.Body.Close()
		Amount := 0.01
		info := result["data"].(map[string]interface{})
		card := info["card"].(map[string]interface{})
		customer := info["customer"].(map[string]interface{})
		log.Println(info)

		msg := events.CreditWallet{
			WalletID: req.FromID,
			Amount:   Amount,
		}
		f.EventEmitter.Emit(&msg, "auth")
		newT.Amount = fmt.Sprintf("%v", info["amount"].(float64)) 
		newT.CardLastFourDigit = card["last_4digits"].(string)
		newT.CardType = card["type"].(string)
		newT.Currency = info["currency"].(string)
		newT.FlwRef = info["currency"].(string)
		newT.Message = req.InnerMemo
		newT.PaymentType = info["payment_type"].(string)
		newT.TransType = 2
		newT.Customer  = customer["email"].(string)
		newT.Status = info["status"].(string)
		newT.TxRef = info["tx_ref"].(string)
		newT.WalletId = req.FromID
		err = f.Db.CreateTransaction(&newT)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument")
		}
		return &v1.NewTransactionRes{
			ID: result["status"].(string),
		}, nil

	case 1:
		sReqURL, _ := url.Parse(fmt.Sprintf("http://localhost:9000/v1/wallet/%v", req.FromID))

		sender := &http.Request{
			Method: "GET",
			URL:    sReqURL,
			Header: map[string][]string{
				"Content-Type":  {"application/json; charset=UTF-8"},
				"Authorization": {authorization},
			},
		}
		res, err := HttpReq(sender)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "Invalid Argument")
		}

		var result map[string]interface{}
		json.NewDecoder(res.Body).Decode(&result)
		// close response body
		res.Body.Close()
		bal := result["accountBal"]
		log.Println(result)
		//bal, ok := result["AccountBal"]
		//if !ok {
		//	//return
		//}
		if bal.(float64) <= req.Amount {
			return nil, status.Error(codes.InvalidArgument, "Low balance")
		}
		msg1 := events.CreditWallet{
			WalletID: req.FromID,
			Amount:   -req.Amount,
		}
		msg2 := events.CreditWallet{
			WalletID: req.ToID,
			Amount:   req.Amount,
		}
		f.EventEmitter.Emit(&msg1, "auth")
		f.EventEmitter.Emit(&msg2, "auth")
		//amt := fmt.Sprintf("%f", req.Amount)

		return &v1.NewTransactionRes{
			ID: "xxx-id",
		}, nil
	}
	return nil, status.Error(codes.InvalidArgument, "Invalid transaction type")

}

func (f *flairsTransactionServer) GetMyTransactions(ctx context.Context, req *v1.GetMyTransactionsRequest) (*v1.TransactionsResponse, error) {
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
	/** read response body
	data, _ := ioutil.ReadAll(res.Body)

	// close response body
	res.Body.Close()

	// print response status and body
	log.Printf("status: %d\n", res.StatusCode)
	log.Printf("body: %s\n", string(data))
	return data, err
	**/
}
