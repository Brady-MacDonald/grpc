package main

import "sync"

type Store struct {
	mu   sync.RWMutex
	urls map[string]string // shortCode -> longURL
}

// NewStore: initializes a new in-memory store for short to long URL mappings
func NewStore() *Store {
	return &Store{
		urls: make(map[string]string),
	}
}

// Save the shortened url mapping
// Return false if unable to save due to it already existing
func (s *Store) Save(code, longURL string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.urls[code]; exists {
		return false
	}

	s.urls[code] = longURL
	return true
}

// Retrieve the previously shortened URL
func (s *Store) Get(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.urls[code]
	return url, ok
}
