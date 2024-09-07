package main

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var producer sarama.SyncProducer

func main() {
	var err error
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err = sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka producer: %v", err)
	}
	defer producer.Close()

	router := gin.Default()

	router.POST("/create", func(c *gin.Context) {
		message := "Order created successfully"
		_, _, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: "order-topic",
			Value: sarama.StringEncoder(message),
		})
		if err != nil {
			log.Printf("Failed to send message to Kafka: %v", err)
		}
		c.JSON(http.StatusOK, gin.H{"message": message})
	})

	router.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Order list retrieved successfully"})
	})

	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Order service failed to start: %v", err)
	}
}
