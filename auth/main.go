package main

import (
	"auth/grpc/authserver"
	redisconn "auth/redis"
	"auth/rest"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	event_amqp "shared/events/amqp"
)

func main() {

	go authserver.Start()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Could not establish amqp connection : " + err.Error())
	}
	eventEmitter, err := event_amqp.NewAMQPEventEmitter(conn, "auth")
	if err != nil {
		log.Fatalf("count not declare an exchange %v", err)
	}
	defer conn.Close()

	redisPool := redisconn.NewPool("localhost:6379")
	redisConn := redisPool.Get()

	defer redisConn.Close()

	r := rest.ServerRoute(eventEmitter, redisConn)
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", r))

}
