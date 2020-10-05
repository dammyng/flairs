package rest

import (
	"auth/grpc/authclient"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"shared/events"
	amqp "shared/events/amqp"

	"golang.org/x/crypto/bcrypt"

	"shared/helper"
	"shared/models/appuser"
	"strings"

	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

// ServiceHandler represent routes dependencies
type ServiceHandler struct {
	EventEmitter amqp.EventEmitter
	RedisConn    redis.Conn
}

// NewServiceHandler : Service handler constructor
func NewServiceHandler(eventEmitter amqp.EventEmitter, redisConn redis.Conn) ServiceHandler {
	return ServiceHandler{
		EventEmitter: eventEmitter,
		RedisConn:    redisConn,
	}
}

//Register a new user
func (serviceHandler ServiceHandler) Register(w http.ResponseWriter, r *http.Request) {
	var payload EmailPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		msg := InvalidRequest
		fmt.Println(msg)
		helper.DisplayAppError(w, err, msg, http.StatusUnprocessableEntity)
		return
	}
	payload.Email = strings.ToLower(payload.Email)

	// Find user
	u, err := authclient.GetAUser(&appuser.UserArg{UserPayload: &appuser.User{Email: payload.Email}}); 
	if u != nil{
		helper.DisplayAppError(
			w,
			fmt.Errorf("Duplicate entry Error"),
			"A Flairs account with this email already exists. Please try logging in.",
			http.StatusUnauthorized,
		)
		return
	}
	if u == nil && err.Error() != "rpc error: code = Unknown desc = record not found" {
		helper.DisplayAppError(
			w, err,
			ProcessingRequestError,
			http.StatusInternalServerError,
		)
		return
	}

	

	// User was not found so we create a new user instance
	ID := uuid.NewV4()
	tempPass := helper.RandString(20)

	hashedPass, err := getPasswordHash(tempPass)
	if err != nil {
		msg := ProcessingRequestError
		helper.DisplayAppError(w, err, msg, http.StatusInternalServerError)
	}
	user := appuser.User{
		ID:       ID.String(),
		Email:    payload.Email,
		Password: hashedPass,
	}

	// Check for referral in link
	ref := r.URL.Query().Get("ref")

	refcode := helper.RandString(15)

	user.RefCode = refcode
	if len(ref) > 0 {

		user.Referrer = ref
	}
	if _, err := authclient.AddNewUser(&appuser.UserArg{UserPayload: &user}); err != nil{
		helper.DisplayAppError(w, err, UserCreationError, http.StatusBadRequest)
		return
	}
	// Send user the message token to verify email
	// Do it with Rabbit

	data := helper.HttpResponse{
		Message: UserCreationSuccessful,
		Code:    http.StatusCreated,
	}
	helper.WriteJsonResponse(w, data, http.StatusOK)}

func getPasswordHash(password string) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return b, err
}

//Register a new user
func (serviceHandler ServiceHandler) Login(w http.ResponseWriter, r *http.Request) {
	var payload EmailPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		msg := InvalidRequest
		fmt.Println(msg)
		return
	}
	payload.Email = strings.ToLower(payload.Email)

	// Find user
	//if _, err = serviceHandler.DbHandler.FindUser()
	// if user exist return a duplicate error
	// return

	// Since user is new
	// Generate uer object - password, referral and referallink

	// Save the new user
	// Return

	// Send user the message token to verify email
	// Do it with Rabbit

	w.Write([]byte(""))
}

func (serviceHandler ServiceHandler) Find(w http.ResponseWriter, r *http.Request) {

	responce, err := authclient.Connect()

	if err != nil {
		msg := InvalidRequest
		helper.DisplayAppError(w, err, msg, http.StatusUnprocessableEntity)
	}
	helper.WriteJsonResponse(w, responce, http.StatusOK)
	ID := uuid.NewV4()

	msg := events.UserCreatedEvent{
		ID:    hex.EncodeToString(ID.Bytes()),
		Host:  "http://localhost:15672/",
		Email: "user.Email",
		Token: "Token",
	}

	serviceHandler.EventEmitter.Emit(&msg, "auth")
}
