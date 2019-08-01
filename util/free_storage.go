package gdsutil

// FreeStorage struct
type FreeStorage struct {
	Mutex
	Store []*interface{}
}

// Empty func
func (fs *FreeStorage) Empty() bool {
	if fs.IsConcurrent() {
		fs.RWMutex.Lock()
		defer fs.RWMutex.RUnlock()
	}
	return len(fs.Store) == 0
}

// EmptyWithoutRUnlock func
func (fs *FreeStorage) EmptyWithoutRUnlock() bool {
	if fs.IsConcurrent() {
		fs.RWMutex.RLock()
	}
	return len(fs.Store) == 0
}

// EmptyWithoutUnlock func
func (fs *FreeStorage) EmptyWithoutUnlock() bool {
	if fs.IsConcurrent() {
		fs.RWMutex.Lock()
	}
	return len(fs.Store) == 0
}

// Size func
func (fs *FreeStorage) Size() uint {
	if fs.IsConcurrent() {
		fs.RWMutex.RLock()
		defer fs.RWMutex.RUnlock()
	}
	return uint(len(fs.Store))
}

// SizeWithoutRUnlock func
func (fs *FreeStorage) SizeWithoutRUnlock() uint {
	if fs.IsConcurrent() {
		fs.RWMutex.RLock()
	}
	return uint(len(fs.Store))
}

// SizeWithoutUnlock func
func (fs *FreeStorage) SizeWithoutUnlock() uint {
	if fs.IsConcurrent() {
		fs.RWMutex.Lock()
	}
	return uint(len(fs.Store))
}

// SizeWithoutConcurrency func
func (fs *FreeStorage) SizeWithoutConcurrency() uint {
	return uint(len(fs.Store))
}

// Clear func
func (fs *FreeStorage) Clear() {
	if fs.IsConcurrent() {
		fs.RWMutex.Lock()
		defer fs.RWMutex.Unlock()
	}
	fs.Store = nil
}
