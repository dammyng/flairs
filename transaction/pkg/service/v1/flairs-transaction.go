package v1

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	v1internals "transaction/internals/v1"
	v1 "transaction/pkg/api/v1"
	"shared/events"
	amqp "shared/events/amqp"

	"github.com/dgrijalva/jwt-go"
)

const (
	apiVersion = "v1"
)

type flairsTransactionServer struct {
	Db v1internals.DatabaseHandler
	EventEmitter amqp.EventEmitter

}

// Claims jwt custom authentication claims
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// NewflairsTransactionServer creates ToDo service
func NewflairsTransactionServer(db v1internals.DatabaseHandler,eventEmitter amqp.EventEmitter) v1.FlairsTransactionServiceServer {
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
	log.Println(req.T_ID)
	reqURL, _ := url.Parse(fmt.Sprintf("https://api.flutterwave.com/v3/transactions/%v/verify", req.T_ID))
	//reqURL, _ := url.Parse(fmt.Sprintf("https://api.flutterwave.com/v3/transactions/%v/verify", req.T_ID))

	// create request body
	//bodyContent := fmt.Sprintf(``)

	//reqBody := ioutil.NopCloser(strings.NewReader(bodyContent))

	flutterReq := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Content-Type":  {"application/json; charset=UTF-8"},
			"Authorization": {"Bearer " + os.Getenv("FlutterSecret")},
		},
		//Body: reqBody,
	}
	HttpReq(flutterReq)

	msg := events.CreditWallet{
		URL: "http://localhost:9000/v1/wallet/transact/" + req.WalletID,
		//	UserID: user.ID,
		//	Token: tokenString,
	}
	f.EventEmitter.Emit(&msg, "auth")

	return nil, nil
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

func HttpReq(req *http.Request) {

	// send an HTTP request using `req` object
	res, err := http.DefaultClient.Do(req)

	// check for response error
	if err != nil {
		log.Fatal("Error:", err)
	}
	// read response body
	data, _ := ioutil.ReadAll(res.Body)

	// close response body
	res.Body.Close()

	// print response status and body
	log.Printf("status: %d\n", res.StatusCode)
	log.Printf("body: %s\n", string(data))
}
