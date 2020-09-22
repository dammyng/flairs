package rest

import (
	"github.com/gorilla/mux"
	"flairs/auth/libs/persistence"

)

//ServerRoute handles HTTP traffic
func ServerRoute(dbHandler persistence.DatabaseHandler) *mux.Router{
	handler := NewServiceHandler(dbHandler)
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.Methods("POST").Path("/register").HandlerFunc(handler.Register)

	return r
}