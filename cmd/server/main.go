package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/varedis/learn-pub-sub-starter/internal/pubsub"
	"github.com/varedis/learn-pub-sub-starter/internal/routing"
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

	publishCh, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not create channel: %v", err)
	}

	err = pubsub.PublishJSON(
		publishCh,
		routing.ExchangePerilDirect,
		routing.PauseKey,
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		log.Printf("could not publish state: %v", err)
	}
	fmt.Println("Pause message sent!")

}
