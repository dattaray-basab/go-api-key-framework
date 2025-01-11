package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateAPIKey() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("Failed to generate API key: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
