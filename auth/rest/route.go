package rest

import (
	"github.com/gorilla/mux"
	amqp "shared/events/amqp"
	"github.com/gomodule/redigo/redis"


)

//ServerRoute handles HTTP traffic
func ServerRoute(eventEmitter amqp.EventEmitter, redisConn redis.Conn ) *mux.Router{
	handler := NewServiceHandler(eventEmitter, redisConn)
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.Methods("GET").Path("/users").HandlerFunc(handler.AllUsers)
	authRouter.Methods("POST").Path("/register").HandlerFunc(handler.Register)

	return r
}