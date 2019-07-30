package golangds

import "sync"

// Mutex struct
type Mutex struct {
	mutex struct {
		sync.RWMutex
		isConcurrent bool
	}
}

// EnableConcurrency func
func (m *Mutex) EnableConcurrency() {
	m.mutex.isConcurrent = true
}

type array struct {
	Mutex
	Store []*interface{}
}

// Empty func
func (a *array) Empty() bool {
	if a.mutex.isConcurrent {
		a.mutex.RLock()
		defer a.mutex.RUnlock()
	}
	return len(a.Store) == 0
}

func (a *array) emptyWithoutRUnlock() bool {
	if a.mutex.isConcurrent {
		a.mutex.RLock()
	}
	return len(a.Store) == 0
}

func (a *array) emptyWithoutUnlock() bool {
	if a.mutex.isConcurrent {
		a.mutex.Lock()
	}
	return len(a.Store) == 0
}

// Size func
func (a *array) Size() uint {
	if a.mutex.isConcurrent {
		a.mutex.RLock()
		defer a.mutex.RUnlock()
	}
	return uint(len(a.Store))
}

func (a *array) sizeWithoutRUnlock() uint {
	if a.mutex.isConcurrent {
		a.mutex.RLock()
	}
	return uint(len(a.Store))
}

func (a *array) sizeWithoutUnlock() uint {
	if a.mutex.isConcurrent {
		a.mutex.Lock()
	}
	return uint(len(a.Store))
}

func (a *array) sizeWithoutConcurrency() uint {
	return uint(len(a.Store))
}

// Clear func
func (a *array) Clear() {
	if a.mutex.isConcurrent {
		a.mutex.Lock()
		defer a.mutex.Unlock()
	}
	a.Store = nil
}

type storage struct {
	Mutex
	store []*interface{}
}

// Empty func
func (s *storage) Empty() bool {
	if s.mutex.isConcurrent {
		s.mutex.RLock()
		defer s.mutex.RUnlock()
	}
	return len(s.store) == 0
}

func (s *storage) emptyWithoutRUnlock() bool {
	if s.mutex.isConcurrent {
		s.mutex.RLock()
	}
	return len(s.store) == 0
}

func (s *storage) emptyWithoutUnlock() bool {
	if s.mutex.isConcurrent {
		s.mutex.Lock()
	}
	return len(s.store) == 0
}

// Size func
func (s *storage) Size() uint {
	if s.mutex.isConcurrent {
		s.mutex.RLock()
		defer s.mutex.RUnlock()
	}
	return uint(len(s.store))
}

func (s *storage) sizeWithoutRUnlock() uint {
	if s.mutex.isConcurrent {
		s.mutex.RLock()
	}
	return uint(len(s.store))
}

func (s *storage) sizeWithoutUnlock() uint {
	if s.mutex.isConcurrent {
		s.mutex.Lock()
	}
	return uint(len(s.store))
}

func (s *storage) sizeWithoutConcurrency() uint {
	return uint(len(s.store))
}

// Clear func
func (s *storage) Clear() {
	if s.mutex.isConcurrent {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	s.store = nil
}
