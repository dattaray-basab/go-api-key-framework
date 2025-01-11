package storage

import "sync"

type KeyMetadata struct {
	UserID string
	Active bool
}

type KeyStore struct {
	mu   sync.RWMutex
	keys map[string]KeyMetadata
}

func NewKeyStore() *KeyStore {
	return &KeyStore{keys: make(map[string]KeyMetadata)}
}

func (s *KeyStore) AddKey(userID, apiKey string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keys[apiKey] = KeyMetadata{UserID: userID, Active: true}
}

func (s *KeyStore) RevokeKey(apiKey string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.keys[apiKey]; exists {
		delete(s.keys, apiKey)
		return true
	}
	return false
}

func (s *KeyStore) IsValid(apiKey string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	meta, exists := s.keys[apiKey]
	return exists && meta.Active
}
