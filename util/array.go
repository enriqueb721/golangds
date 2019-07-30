package gdsutil

// Array struct
type Array struct {
	Mutex
	Store []*interface{}
}

// Empty func
func (a *Array) Empty() bool {
	if a.IsConcurrent() {
		a.RWMutex.Lock()
		defer a.RWMutex.RUnlock()
	}
	return len(a.Store) == 0
}

// EmptyWithoutRUnlock func
func (a *Array) EmptyWithoutRUnlock() bool {
	if a.IsConcurrent() {
		a.RWMutex.RLock()
	}
	return len(a.Store) == 0
}

// EmptyWithoutUnlock func
func (a *Array) EmptyWithoutUnlock() bool {
	if a.IsConcurrent() {
		a.RWMutex.Lock()
	}
	return len(a.Store) == 0
}

// Size func
func (a *Array) Size() uint {
	if a.IsConcurrent() {
		a.RWMutex.RLock()
		defer a.RWMutex.RUnlock()
	}
	return uint(len(a.Store))
}

// SizeWithoutRUnlock func
func (a *Array) SizeWithoutRUnlock() uint {
	if a.IsConcurrent() {
		a.RWMutex.RLock()
	}
	return uint(len(a.Store))
}

// SizeWithoutUnlock func
func (a *Array) SizeWithoutUnlock() uint {
	if a.IsConcurrent() {
		a.RWMutex.Lock()
	}
	return uint(len(a.Store))
}

// SizeWithoutConcurrency func
func (a *Array) SizeWithoutConcurrency() uint {
	return uint(len(a.Store))
}

// Clear func
func (a *Array) Clear() {
	if a.IsConcurrent() {
		a.RWMutex.Lock()
		defer a.RWMutex.Unlock()
	}
	a.Store = nil
}
