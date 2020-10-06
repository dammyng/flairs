package rest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"shared/helper"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var err error

func DecodeJwt(token string, claims *Claims) error {
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTKey")), nil
	})
	return err
}

func IsAuthenticatedMiddleWare(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	headers := helper.GetHeaders(r, "access-token")
	sessId := headers["access-token"]
	if sessId == nil {

		helper.WriteJsonResponse(w, helper.ErrorObj{
			Error:   errors.New("missing request header value").Error(),
			Message: "access-token header missing",
		}, http.StatusForbidden)
		return
	}

	claims := &Claims{}

	err = DecodeJwt(sessId.(string), claims)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			json.NewEncoder(w).Encode(helper.ErrorObj{Message: err.Error(), Error: err.Error()})

			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if claims.Valid() != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Printf("user: %v\n", claims.UserId)
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Millisecond)
	ctx = context.WithValue(ctx, "user", claims.UserId)
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)
	defer cancel()
}