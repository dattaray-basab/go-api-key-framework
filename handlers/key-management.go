package handlers

import (
	"net/http"

	"github.com/dattaray-basab/go-api-key-management/utils"
	"github.com/gin-gonic/gin"

	"github.com/dattaray-basab/go-api-key-management/storage"
)

var store = storage.NewKeyStore()

// GenerateKey creates a new API key for a user.
func GenerateKey(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	apiKey := utils.GenerateAPIKey()
	store.AddKey(req.UserID, apiKey)

	c.JSON(http.StatusOK, gin.H{"api_key": apiKey})
}

// RevokeKey revokes an API key.
func RevokeKey(c *gin.Context) {
	var req struct {
		APIKey string `json:"api_key" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "api_key is required"})
		return
	}

	if store.RevokeKey(req.APIKey) {
		c.JSON(http.StatusOK, gin.H{"message": "API key revoked"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "API key not found"})
	}
}

// ValidateKey checks if an API key is valid.
func ValidateKey(c *gin.Context) {
	apiKey := c.Query("api_key")
	if apiKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "api_key is required"})
		return
	}

	if store.IsValid(apiKey) {
		c.JSON(http.StatusOK, gin.H{"message": "API key is valid"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
	}
}
