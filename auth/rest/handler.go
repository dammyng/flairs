package rest

import (
	"encoding/json"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"shared/helper"
	"shared/events"
	"auth/grpc/authclient"
	uuid "github.com/satori/go.uuid"

	amqp "shared/events/amqp"
	"github.com/gomodule/redigo/redis"

)

// ServiceHandler represent routes dependencies
type ServiceHandler struct {
	EventEmitter amqp.EventEmitter
	RedisConn redis.Conn

}

// NewServiceHandler : Service handler constructor
func NewServiceHandler(eventEmitter amqp.EventEmitter,redisConn redis.Conn) ServiceHandler {
	return ServiceHandler{
		EventEmitter: eventEmitter,
		RedisConn : redisConn,
	}
}

//Register a new user
func (serviceHandler ServiceHandler) Register(w http.ResponseWriter, r *http.Request) {
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

func (serviceHandler ServiceHandler) AllUsers(w http.ResponseWriter, r *http.Request) {
	
	responce, err := authclient.Connect()
	
	if err != nil {
		msg := InvalidRequest
		helper.DisplayAppError(w, err, msg, http.StatusUnprocessableEntity)
	}
	helper.WriteJsonResponse(w, responce, http.StatusOK)
	ID := uuid.NewV4()

	msg:= events. UserCreatedEvent{
		ID: hex.EncodeToString(ID.Bytes()),
		Host: "http://localhost:15672/",
		Email: "user.Email",
		Token: "Token",
	}

	serviceHandler.EventEmitter.Emit(&msg, "auth")
}

