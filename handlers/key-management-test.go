package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// SetupRouter initializes the Gin router for testing.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/generate", GenerateKey)
	r.POST("/revoke", RevokeKey)
	r.GET("/validate", ValidateKey)
	return r
}

func TestAPIKeyFlow(t *testing.T) {
	router := SetupRouter()

	// Step 1: Generate API Key
	generatePayload := map[string]string{"user_id": "test_user"}
	generateBody, _ := json.Marshal(generatePayload)

	req := httptest.NewRequest(http.MethodPost, "/generate", bytes.NewBuffer(generateBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Assert the status code and response
	assert.Equal(t, http.StatusOK, w.Code)

	var generateResponse map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &generateResponse)
	assert.NoError(t, err)

	apiKey, exists := generateResponse["api_key"]
	assert.True(t, exists)
	assert.NotEmpty(t, apiKey)

	// Step 2: Validate the generated API Key
	req = httptest.NewRequest(http.MethodGet, "/validate?api_key="+apiKey, nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "API key is valid\n", w.Body.String())

	// Step 3: Revoke the API Key
	revokePayload := map[string]string{"api_key": apiKey}
	revokeBody, _ := json.Marshal(revokePayload)

	req = httptest.NewRequest(http.MethodPost, "/revoke", bytes.NewBuffer(revokeBody))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "API key revoked\n", w.Body.String())

	// Step 4: Validate the revoked API Key
	req = httptest.NewRequest(http.MethodGet, "/validate?api_key="+apiKey, nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"error\":\"Invalid API key\"}\n", w.Body.String())
}
