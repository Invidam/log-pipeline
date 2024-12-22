package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {

	log.Println("Starting Kafka consumer...")
	time.Sleep(5 * time.Second)
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
			log.Printf("Error reading message: %v", err)
			// Wait for a few seconds before retrying
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Printf("Message: %s\n", string(msg.Value))
	}
}
