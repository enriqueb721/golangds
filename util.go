package golangds

type storage struct {
	store []*interface{}
}

// Empty func
func (s *storage) Empty() bool {
	return len(s.store) == 0
}

// Size func
func (s *storage) Size() uint {
	return uint(len(s.store))
}

// Clear func
func (s *storage) Clear() {
	s.store = nil
}
