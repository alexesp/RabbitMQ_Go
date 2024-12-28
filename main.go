package main

import (
	"fmt"

	"github.com/alexesp/RabbitMQ_Go.git/internal/rabbitmq"
)

type App struct {
	Rmq *rabbitmq.RabbitMQ
}

func Run() error {
	fmt.Println("Go Rabbitmq")

	rmq := rabbitmq.NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error Setting Up our application")
		fmt.Println(err)
	}
}
