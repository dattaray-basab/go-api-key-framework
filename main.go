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


// to test using curl
// curl -X POST http://localhost:8080/generate -H "Content-Type: application/json" -d '{"user_id": "test_user"}'
// curl -X GET "http://localhost:8080/validate?api_key=<YOUR_API_KEY>"
// curl -X POST http://localhost:8080/revoke -H "Content-Type: application/json" -d '{"api_key": "<YOUR_API_KEY>"}'

// To test the API key management system, you can use the following curl commands:
// Actual STEPS shown below:
// 1.
// bd@Basabs-MBP go-api-key-framework % curl -X POST http://localhost:8080/generate -H "Content-Type: application/json" -d '{"user_id": "test_user"}'

// {"api_key":"NlINIdDZfT0y6RxKRVGC1cYWWK9XMPO5tSFdPaYOC_A="}% 

// 2.
// bd@Basabs-MBP go-api-key-framework % curl -X GET "http://localhost:8080/validate?api_key=NlINIdDZfT0y6RxKRVGC1cYWWK9XMPO5tSFdPaYOC_A="
// {"message":"API key is valid"}%    

// 3.
// bd@Basabs-MBP go-api-key-framework % curl -X POST http://localhost:8080/revoke -H "Content-Type: application/json" -d '{"api_key": "NlINIdDZfT0y6RxKRVGC1cYWWK9XMPO5tSFdPaYOC_A="}'
// {"message":"API key revoked"}%   

