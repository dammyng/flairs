package v1

import (
	"google.golang.org/grpc/metadata"
	"context"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	v1internals "transaction/internals/v1"
	v1 "transaction/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

type flairsTransactionServer struct {
	Db v1internals.DatabaseHandler
}

// Claims jwt custom authentication claims
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// NewflairsTransactionServer creates ToDo service
func NewflairsTransactionServer(db v1internals.DatabaseHandler) v1.Transaction {
	return &flairsTransactionServer{Db: db}
}

// DecodeJwt - decodes JWT token from request
func DecodeJwt(token string, claims *Claims) error {
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secrek_key"), nil
	})
	return err
}

   //{status: "successful", customer: {…}, transaction_id: 1695241, tx_ref: "hooli-tx-1920bbtyt", flw_ref: "FLW-MOCK-6439899760b3449a2db802decd80594f", …}
//amount: 100
//currency: "NGN"
//customer: {name: "yemi desola", email: "user@gmail.com", phone_number: "08102909304"}
//flw_ref: "FLW-MOCK-6439899760b3449a2db802decd80594f"
//status: "successful"
//transaction_id: 1695241
//tx_ref: "hooli-tx-1920bbtyt"