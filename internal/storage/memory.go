package storage

import (
	"sync"
	"uptime-monitor/internal/check"
)

type MemoryStorage struct {
	mu     sync.Mutex
	checks map[string]check.Check
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		checks: make(map[string]check.Check),
	}
}

func (s *MemoryStorage) Save(c check.Check) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.checks[c.ID] = c
}

func (s *MemoryStorage) GetAll() []check.Check {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]check.Check, 0, len(s.checks))
	for _, c := range s.checks {
		result = append(result, c)
	}

	return result
}