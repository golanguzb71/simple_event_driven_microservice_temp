package main

import (
	"github.com/IBM/sarama"
	"log"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("order-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start Kafka partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		log.Printf("Received message: %s", string(message.Value))
	}
}
