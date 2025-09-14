package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	godotenv.Load()
	fmt.Println("Starting Peril server...")

	rabbitMQConnectionString := os.Getenv("RABBIT_MQ_CONNECTION")
	if rabbitMQConnectionString == "" {
		log.Fatal("RABBIT_MQ_CONNECTION is not set")
	}

	conn, err := amqp.Dial(rabbitMQConnectionString)
	if err != nil {
		log.Fatalf("could not connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connection to RabbitMQ successful")

	// Wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	fmt.Println("Shutting down Peril server...")
}
