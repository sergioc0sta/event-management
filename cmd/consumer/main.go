package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sergioc0sta/event-management/pkg/queue"
)

func main(){
	fmt.Print("Start msg queue")
 ch, err := queue.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go queue.Consumer(ch, msgs)

	for msg := range msgs {
		fmt.Println("Messange: ", string(msg.Body))
		msg.Ack(false)
	}


}
