package queue

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	return channel, nil
}

func Consumer(ch *amqp.Channel, msgOut chan<- amqp.Delivery) error {
	_, err := ch.QueueDeclare(
		"myQueue",
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		panic(err)
	}

	msgs, err := ch.Consume(
		"myQueue",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		msgOut <- msg
	}
	return nil
}
