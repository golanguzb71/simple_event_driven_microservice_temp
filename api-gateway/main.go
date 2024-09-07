package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	router.Any("/user/*any", func(c *gin.Context) {
		proxy(c, "http://localhost:8081")
	})

	router.Any("/order/*any", func(c *gin.Context) {
		proxy(c, "http://localhost:8082")
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("API Gateway failed to start: %v", err)
	}
}

func proxy(c *gin.Context, target string) {
	req, err := http.NewRequest(c.Request.Method, target, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header = c.Request.Header

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to connect to service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Writer.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	c.Writer.Header().Set("Content-Length", strconv.Itoa(len(body)))
	c.Writer.WriteHeader(resp.StatusCode)

	c.Writer.Write(body)
}
