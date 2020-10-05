package main

import (
	"auth/grpc/authserver"
	"auth/libs/persistence"
	"auth/libs/setup"
	"auth/rest"
	"log"
	"os"

	redisconn "auth/redis"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"net/http"
	event_amqp "shared/events/amqp"
)

// App represent the module
type App struct {
	Router     *mux.Router
	DbHandler  persistence.DatabaseHandler
	authServer authserver.AuthServer
}

// LoadEnv load env file
func LoadEnv() {
	log.Println("env loading...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// InitDB would start grpc server on TCP port addr
func (a *App) InitDB() {
	LoadEnv()
	dbHandler := persistence.NewMysqlLayer(os.Getenv("DBConfig"))
	dbHandler.Session.Exec(setup.SetTimeZone)
	dbHandler.Session.Exec(setup.CreateDatabase)
	dbHandler.Session.Exec(setup.UseAlphaPlus)
	dbHandler.Session.Exec(setup.CreateUserTable)
	a.DbHandler = dbHandler
}

// StartGRPC would start grpc server on TCP port addr
func (a *App) StartGRPC() {
	authS := authserver.NewAuthServer(a.DbHandler)

	go authserver.Start(authS, os.Getenv("GRPCPort"))
	a.authServer = authS
}

// InitHandler set up the http handler for the requests
func (a *App) InitHandler() {
	// Handler require AMQP so itt is initialized within
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		log.Fatal("Could not establish amqp connection : " + err.Error())
	}
	eventEmitter, err := event_amqp.NewAMQPEventEmitter(conn, "auth")
	if err != nil {
		log.Fatalf("count not declare an exchange %v", err)
	}
	defer conn.Close()

	// Handler require Redis so it is initialized within
	redisPool := redisconn.NewPool(os.Getenv("REDIS_URL"))
	redisConn := redisPool.Get()

	//defer redisConn.Close()

	r := rest.ServerRouter(eventEmitter, redisConn)

	a.Router = r
}

// StartHTTP would start http server on port addr
func (a *App) StartHTTP() {
	log.Fatal(http.ListenAndServe(os.Getenv("HTTPPort"), a.Router))
}
