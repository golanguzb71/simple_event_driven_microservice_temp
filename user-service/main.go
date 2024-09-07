package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/register", func(c *gin.Context) {
		// Foydalanuvchini ro'yxatdan o'tkazish logikasi
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	router.POST("/login", func(c *gin.Context) {
		// Foydalanuvchini kirish logikasi
		c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
	})

	if err := router.Run(":8081"); err != nil {
		log.Fatalf("User service failed to start: %v", err)
	}
}
