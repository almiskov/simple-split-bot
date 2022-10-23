package store

import (
	"strings"
	"sync"
)

type Store struct {
	mu sync.Mutex
	m  map[int64][]string
}

func New() *Store {
	return &Store{
		m: make(map[int64][]string),
	}
}

func (s *Store) Init(userID int64) {
	s.mu.Lock()
	s.m[userID] = make([]string, 0)
	s.mu.Unlock()
}

func (s *Store) Add(userID int64, text string) {
	s.mu.Lock()
	s.m[userID] = append(s.m[userID], text)
	s.mu.Unlock()
}

func (s *Store) Get(userID int64) (string, bool) {
	msgs, ok := s.m[userID]
	return strings.Join(msgs, "\n"), ok
}

func (s *Store) Clear(userID int64) {
	s.mu.Lock()
	delete(s.m, userID)
	s.mu.Unlock()
}
