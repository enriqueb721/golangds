package gdsutil

import "sync"

// Mutex struct
type Mutex struct {
	RWMutex      sync.RWMutex
	isConcurrent bool
}

// EnableConcurrency func
func (m *Mutex) EnableConcurrency() {
	m.isConcurrent = true
}

// IsConcurrent func
func (m *Mutex) IsConcurrent() bool {
	return m.isConcurrent
}
