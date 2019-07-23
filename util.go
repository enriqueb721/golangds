package golangds

import "sync"

// storage struct
type storage struct {
	mutex sync.RWMutex
	store []*interface{}

	isConcurrent bool
}

// EnableConcurrent func
func (s *storage) EnableConcurrency() {
	s.isConcurrent = true
}

// Empty func
func (s *storage) Empty() bool {
	if s.isConcurrent {
		s.mutex.RLock()
		defer s.mutex.RUnlock()
	}
	return len(s.store) == 0
}

// Size func
func (s *storage) Size() uint {
	if s.isConcurrent {
		s.mutex.RLock()
		defer s.mutex.RUnlock()
	}
	return uint(len(s.store))
}

// Clear func
func (s *storage) Clear() {
	if s.isConcurrent {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	s.store = nil
}
