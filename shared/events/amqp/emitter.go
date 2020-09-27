// Emitter
package amqp

import (
	"encoding/json"
	"fmt"
	"shared/events"

	"github.com/streadway/amqp"
)

type EventEmitter interface{
	Emit(event events.Event, exchange string) error
}

type ampqEventEmitter struct{
	connection *amqp.Connection
}

func (a *ampqEventEmitter) setup(exchange string) error {
	channel, err := a.connection.Channel()
	if err != nil{
		return err
	}
	defer channel.Close()

	return channel.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
}

func NewAMQPEventEmitter(conn *amqp.Connection, exchange string) (EventEmitter, error) {
	emitter := &ampqEventEmitter{
		connection: conn,
	}

	err := emitter.setup(exchange)

	if err != nil {
		return nil, err
	}
	return emitter, nil
}

func (a *ampqEventEmitter)Emit(event events.Event, exchange string) error  {
	jsonData, err := json.Marshal(event)

	if err != nil{
		return nil
	}

	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		Headers:amqp.Table{
			"x-event-name":event.EventName(),
		},
		Body:jsonData,
		ContentType:"application/json",
	}

	fmt.Printf("emit msg %s ", msg)
	return channel.Publish(exchange, event.EventName(), false, false, msg)
}