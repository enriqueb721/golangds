package gdsutil

// Storage struct
type Storage struct {
	Mutex
	store []*interface{}
}

func (s *Storage) Store() []*interface{} {
	return s.store
}

func (s *Storage) SetStore(store []*interface{}) {
	s.store = store
}

// Empty func
func (s *Storage) Empty() bool {
	if s.IsConcurrent() {
		s.RWMutex.RLock()
		defer s.RWMutex.RUnlock()
	}
	return len(s.store) == 0
}

func (s *Storage) EmptyWithoutRUnlock() bool {
	if s.IsConcurrent() {
		s.RWMutex.RLock()
	}
	return len(s.store) == 0
}

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

func (s *Storage) SizeWithoutRUnlock() uint {
	if s.IsConcurrent() {
		s.RWMutex.RLock()
	}
	return uint(len(s.store))
}

func (s *Storage) SizeWithoutUnlock() uint {
	if s.IsConcurrent() {
		s.RWMutex.Lock()
	}
	return uint(len(s.store))
}

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
