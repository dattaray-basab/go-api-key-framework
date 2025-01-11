package main

import (
	"log"

	"github.com/dattaray-basab/go-api-key-management/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the in-memory key store (or replace with a persistent store if needed)

	// Initialize the Gin router
	router := gin.Default()

	// API Endpoints
	router.POST("/generate", handlers.GenerateKey)
	router.POST("/revoke", handlers.RevokeKey)
	router.GET("/validate", handlers.ValidateKey)

	// Start the server
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
