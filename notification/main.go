package main

import (
	"fmt"
	"log"

	event_amqp "shared/events/amqp"
	"shared/events"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("could not establish ampy: ", err.Error())
	}
	defer conn.Close()
	eventListener, err := event_amqp.NewAMQPEventListener(conn, "auth")
	go ProcessEvents(eventListener)
	c := make(chan int)
	<-c
}

func ProcessEvents(eventListener events.EventListener) error {
	received, errors, err := eventListener.Listen("auth", "user.created")
	if err != nil {
		log.Fatalf("event listenner error")
	}

	//templateBox, err := rice.FindBox("html")
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case evt := <- received:
			fmt.Printf("got event %s " , evt.EventName())
		
			switch e:=evt.(type) {
			case *events.UserCreatedEvent:
				fmt.Println("got here")
			default:
				log.Printf("unknown event: %t", e)
			}
		case err = <-errors:
			log.Printf(" recieved error while processing msg: %s", err)
		}
	}
}
