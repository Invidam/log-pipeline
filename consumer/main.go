package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Define Kafka reader config
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "log-topic",
	})

	fmt.Println("Kafka consumer started...: no-group")

	// Continuously read messages from Kafka
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("failed to read message: %v", err)
		}
		fmt.Printf("Message: %s\n", string(msg.Value))
	}
}
