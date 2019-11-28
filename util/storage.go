package gdsutil

// Storage struct
type Storage struct {
	Mutex
	store []*interface{}
}

// Store func
func (s *Storage) Store() []*interface{} {
	if s.IsConcurrent() {
		s.RWMutex.Lock()
		defer s.RWMutex.RUnlock()
	}
	return s.store
}

// SetStore func
func (s *Storage) SetStore(store []*interface{}) {
	if s.IsConcurrent() {
		s.RWMutex.Lock()
		defer s.RWMutex.RUnlock()
	}
	s.store = store
}

// Empty func
func (s *Storage) Empty() bool {
	if s.IsConcurrent() {
		s.RWMutex.Lock()
		defer s.RWMutex.RUnlock()
	}
	return len(s.store) == 0
}

// EmptyWithoutRUnlock func
func (s *Storage) EmptyWithoutRUnlock() bool {
	if s.IsConcurrent() {
		s.RWMutex.RLock()
	}
	return len(s.store) == 0
}

// EmptyWithoutUnlock func
func (s *Storage) EmptyWithoutUnlock() bool {
	if s.IsConcurrent() {
		s.RWMutex.Lock()
	}
	return len(s.store) == 0
}

// Size func
func (s *Storage) Size() uint {
	if s.IsConcurrent() {
		s.RWMutex.RLock()
		defer s.RWMutex.RUnlock()
	}
	return uint(len(s.store))
}

// SizeWithoutRUnlock func
func (s *Storage) SizeWithoutRUnlock() uint {
	if s.IsConcurrent() {
		s.RWMutex.RLock()
	}
	return uint(len(s.store))
}

// SizeWithoutUnlock func
func (s *Storage) SizeWithoutUnlock() uint {
	if s.IsConcurrent() {
		s.RWMutex.Lock()
	}
	return uint(len(s.store))
}

// SizeWithoutConcurrency func
func (s *Storage) SizeWithoutConcurrency() uint {
	return uint(len(s.store))
}

// Clear func
func (s *Storage) Clear() {
	if s.IsConcurrent() {
		s.RWMutex.Lock()
		defer s.RWMutex.Unlock()
	}
	s.store = nil
}
