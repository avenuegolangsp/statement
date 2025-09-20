package main

import (
	"fmt"
	"log"
	"time"

	"github.com/avenue-golang/statement/pkg/db"
	"github.com/avenue-golang/statement/pkg/messaging"
)

func main() {
	// Initialize database connection
	dbConn := db.Init()
	if dbConn.Error != nil {
		log.Fatalf("Failed to connect to database: %v", dbConn.Error)
	}

	// Connect to RabbitMQ
	conn := messaging.Connect()
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"events", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// Here you would process the event and interact with the database
			// For now, just logging and simulating some work
			time.Sleep(1 * time.Second)
			fmt.Println("Event processed and saved to DB (simulated)")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}