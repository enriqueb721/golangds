package golangds

import "sync"

// storage struct
type storage struct {
	sync.RWMutex
	store []*interface{}
}

// Empty func
func (s *storage) Empty() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.store) == 0
}

// Size func
func (s *storage) Size() uint {
	s.RLock()
	defer s.RUnlock()
	return uint(len(s.store))
}

// Clear func
func (s *storage) Clear() {
	s.Lock()
	defer s.Unlock()
	s.store = nil
}
