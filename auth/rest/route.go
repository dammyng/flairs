package rest

import (
	"github.com/gorilla/mux"

)

//ServerRoute handles HTTP traffic
func ServerRoute() *mux.Router{
	handler := NewServiceHandler()
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.Methods("GET").Path("/users").HandlerFunc(handler.AllUsers)
	authRouter.Methods("POST").Path("/register").HandlerFunc(handler.Register)

	return r
}