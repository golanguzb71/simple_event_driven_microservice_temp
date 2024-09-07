// order-service/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/create", func(c *gin.Context) {
		// Buyurtma yaratish logikasi
		c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
	})

	router.GET("/list", func(c *gin.Context) {
		// Buyurtmalar ro'yxatini olish logikasi
		c.JSON(http.StatusOK, gin.H{"message": "Order list retrieved successfully"})
	})

	if err := router.Run(":8082"); err != nil {
		log.Fatalf("Order service failed to start: %v", err)
	}
}
