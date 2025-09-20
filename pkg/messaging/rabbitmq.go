package messaging

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func Connect() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	fmt.Println("Connected to RabbitMQ successfully!")
	return conn
}
