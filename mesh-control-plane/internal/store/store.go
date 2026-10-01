package store

import "sync"

type State struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewState() *State {
	return &State{data: map[string]string{}}
}

func (s *State) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *State) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.data[key]
	return value, ok
}
