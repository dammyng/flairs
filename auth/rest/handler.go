package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"auth/libs/persistence"
	"shared/helper"
)

// ServiceHandler represent routes dependencies
type ServiceHandler struct {
	DbHandler persistence.DatabaseHandler
}

// NewServiceHandler : Service handler constructor
func NewServiceHandler(dbHandler persistence.DatabaseHandler) ServiceHandler {
	return ServiceHandler{
		DbHandler: dbHandler,
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
	
	data , err := serviceHandler.DbHandler.AllUsers()
	if err != nil {
		msg := InvalidRequest
		helper.DisplayAppError(w, err, msg, http.StatusUnprocessableEntity)
	}
	helper.WriteJsonResponse(w, data, http.StatusOK)
}
