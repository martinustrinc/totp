package services

import (
	"fmt"
	"sync"
)

// SecretKeyStore stores secret keys for multiple users.
type SecretKeyStore struct {
	mu        sync.RWMutex
	userKeys  map[string]string // Map to store user-secret key pairs
	keyToUser map[string]string // Map to store secret key to user mapping
}

// NewSecretKeyStore initializes a new SecretKeyStore.
func NewSecretKeyStore() *SecretKeyStore {
	return &SecretKeyStore{
		userKeys:  make(map[string]string),
		keyToUser: make(map[string]string),
	}
}

// SetSecretKey sets the secret key for the specified user.
func (s *SecretKeyStore) SetSecretKey(user, key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.userKeys[user] = key
	s.keyToUser[key] = user
}

// GetSecretKey retrieves the secret key for the specified user.
func (s *SecretKeyStore) GetSecretKey(user string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	key, ok := s.userKeys[user]
	if !ok {
		return "", fmt.Errorf("secret key not found for user: %s", user)
	}
	return key, nil
}

// GetUserBySecretKey retrieves the user associated with the specified secret key.
func (s *SecretKeyStore) GetUserBySecretKey(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	user, ok := s.keyToUser[key]
	if !ok {
		return "", fmt.Errorf("user not found for secret key")
	}
	return user, nil
}
