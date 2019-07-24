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

func (s *storage) emptyWithoutRUnlock() bool {
	// th1 pop pop
	// th2 top (se queda en line s.isConcurrent) // runtime error
	// arr[1, 2]
	if s.isConcurrent {
		s.mutex.RLock()
	}
	return len(s.store) == 0
}

func (s *storage) emptyWithoutUnlock() bool {
	if s.isConcurrent {
		s.mutex.Lock()
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

func (s *storage) sizeWithoutRUnlock() bool {
	if s.isConcurrent {
		s.mutex.RLock()
	}
	return len(s.store) == 0
}

// Clear func
func (s *storage) Clear() {
	if s.isConcurrent {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	s.store = nil
}

// Setting struct
type Setting struct {
	disableUnlock bool
}
