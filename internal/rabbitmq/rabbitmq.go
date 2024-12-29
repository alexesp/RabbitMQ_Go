package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Service interface {
	Connect() error
	Publish(message string) error
	Consume()
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")
	var err error

	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err

	}
	fmt.Println("Successfully Connected to RabbitMQ")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}
	_, err = r.Channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	return nil
}

func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}

func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "test/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("Successfully published message to queue")

	return nil
}

func (r *RabbitMQ) Consume() {
	msgs, err := r.Channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	for msg := range msgs {
		fmt.Printf("Received Mesage: %s\n", msg.Body)
	}
}
