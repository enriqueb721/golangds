package gdsutil

// MutableStorage struct
type MutableStorage struct {
	Mutex
	Store []*interface{}
}

// Empty func
func (ms *MutableStorage) Empty() bool {
	if ms.IsConcurrent() {
		ms.RWMutex.Lock()
		defer ms.RWMutex.RUnlock()
	}
	return len(ms.Store) == 0
}

// EmptyWithoutRUnlock func
func (ms *MutableStorage) EmptyWithoutRUnlock() bool {
	if ms.IsConcurrent() {
		ms.RWMutex.RLock()
	}
	return len(ms.Store) == 0
}

// EmptyWithoutUnlock func
func (ms *MutableStorage) EmptyWithoutUnlock() bool {
	if ms.IsConcurrent() {
		ms.RWMutex.Lock()
	}
	return len(ms.Store) == 0
}

// Size func
func (ms *MutableStorage) Size() uint {
	if ms.IsConcurrent() {
		ms.RWMutex.RLock()
		defer ms.RWMutex.RUnlock()
	}
	return uint(len(ms.Store))
}

// SizeWithoutRUnlock func
func (ms *MutableStorage) SizeWithoutRUnlock() uint {
	if ms.IsConcurrent() {
		ms.RWMutex.RLock()
	}
	return uint(len(ms.Store))
}

// SizeWithoutUnlock func
func (ms *MutableStorage) SizeWithoutUnlock() uint {
	if ms.IsConcurrent() {
		ms.RWMutex.Lock()
	}
	return uint(len(ms.Store))
}

// SizeWithoutConcurrency func
func (ms *MutableStorage) SizeWithoutConcurrency() uint {
	return uint(len(ms.Store))
}

// Clear func
func (ms *MutableStorage) Clear() {
	if ms.IsConcurrent() {
		ms.RWMutex.Lock()
		defer ms.RWMutex.Unlock()
	}
	ms.Store = nil
}
