package rest

import (
	"github.com/gorilla/mux"
	amqp "shared/events/amqp"
	"github.com/gomodule/redigo/redis"


)

//ServerRouter handles HTTP traffic
func ServerRouter(eventEmitter amqp.EventEmitter, redisConn redis.Conn ) *mux.Router{
	handler := NewServiceHandler(eventEmitter, redisConn)
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.Methods("GET").Path("/user").HandlerFunc(handler.Find)
	authRouter.Methods("POST").Path("/register").HandlerFunc(handler.Register)
	//authRouter.Methods("PUT").Path("/users").HandlerFunc(handler.Update)
	//authRouter.Methods("PUT").Path("/login").HandlerFunc(handler.Login)

	return r
}