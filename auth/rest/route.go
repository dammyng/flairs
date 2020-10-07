package rest

import (
	amqp "shared/events/amqp"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//ServerRouter handles HTTP traffic
func ServerRouter(eventEmitter amqp.EventEmitter, redisConn redis.Conn) *mux.Router {
	handler := NewServiceHandler(eventEmitter, redisConn)
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.Methods("GET").Path("/user").HandlerFunc(handler.Find)
	authRouter.Methods("POST").Path("/register").HandlerFunc(handler.Register)
	//authRouter.Methods("PUT").Path("/users").HandlerFunc(handler.Update)
	//authRouter.Methods("PUT").Path("/login").HandlerFunc(handler.Login)
	authRouter.Methods("GET").Path("/verifyemail/{email}/{token}").HandlerFunc(handler.VerifyEmail)
	authRouter.Methods("POST").Path("/set_password/{email}/{token}").HandlerFunc(handler.SetPassword)
	authRouter.Methods("POST").Path("/login").HandlerFunc(handler.Login)

	acctBase := mux.NewRouter()

	r.PathPrefix("/account").Handler(negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(IsAuthenticatedMiddleWare),
		negroni.NewLogger(),
		negroni.Wrap(acctBase),
	))

	acctRouter := acctBase.PathPrefix("/account").Subrouter()

	acctRouter.Methods("GET", "PATCH").Path("/user").HandlerFunc(handler.UserData)
	acctRouter.Methods("GET", "POST").Path("/card_request").HandlerFunc(handler.CardRequest)
	acctRouter.Methods("GET", "POST").Path("/wallets").HandlerFunc(handler.Wallet)


	return r
}
