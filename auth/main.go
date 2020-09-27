package main

import (
	"log"
	"net/http"
	"auth/rest"
	"auth/grpc/authserver"
	"github.com/streadway/amqp"
	event_amqp "shared/events/amqp"

)

func main() {
	
	go authserver.Start()	

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal("Could not establish amqp connection : "  + err.Error())
	}

	eventEmitter, err := event_amqp.NewAMQPEventEmitter(conn, "auth")
	if err != nil {
		log.Fatalf("count not declare an exchange %v", err)
	}
	defer conn.Close()


	r := rest.ServerRoute(eventEmitter)
	log.Fatal(http.ListenAndServe("0.0.0.0:9000", r))

}
