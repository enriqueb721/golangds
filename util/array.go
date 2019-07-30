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

func (a *Array) EmptyWithoutRUnlock() bool {
	if a.IsConcurrent() {
		a.RWMutex.RLock()
	}
	return len(a.Store) == 0
}

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

func (a *Array) SizeWithoutRUnlock() uint {
	if a.IsConcurrent() {
		a.RWMutex.RLock()
	}
	return uint(len(a.Store))
}

func (a *Array) SizeWithoutUnlock() uint {
	if a.IsConcurrent() {
		a.RWMutex.Lock()
	}
	return uint(len(a.Store))
}

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
