package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/varedis/learn-pub-sub-starter/internal/gamelogic"
	"github.com/varedis/learn-pub-sub-starter/internal/pubsub"
	"github.com/varedis/learn-pub-sub-starter/internal/routing"
)

func main() {
	fmt.Println("Starting Peril client...")

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

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatalf("could not get username: %v", err)
	}

	_, queue, err := pubsub.DeclareAndBind(
		conn,
		routing.ExchangePerilDirect,
		routing.PauseKey+"."+username,
		routing.PauseKey,
		pubsub.SimpleQueueTransient,
	)
	if err != nil {
		log.Fatalf("could not subscribe to pause: %v", err)
	}
	fmt.Printf("Queue %v declared and bound!\n", queue.Name)

	// Wait for ctrl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
