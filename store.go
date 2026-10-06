package main

import "sync"

// Store is an in-memory key-value store
// It is safe for concurrent use: mu guards every access to data
type Store struct {
	mu   sync.Mutex
	data map[string]string
}

// NewStore returns an empty store ready to use
// The map must be initialized here, or a nil map would crash the Put function
func NewStore() *Store {
	return &Store{
		data: map[string]string{},
	}
}

// Get returns the value at the key, ok is false if the key doesn't exist
// which lets callers tell a missing key apart from an empty value
func (s *Store) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[key]
	return value, ok
}

// Puts values under a key, and if a value already exists at the key, it's overwritten
func (s *Store) Put(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

// Deletes removes the key and its values, and does nothing if key doesn't exist
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.data, key)
}
