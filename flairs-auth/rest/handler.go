package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ServiceHandler represent routes dependencies
type ServiceHandler struct {
}

// NewServiceHandler : Service handler constructor
func NewServiceHandler() ServiceHandler {
	return ServiceHandler{}
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
