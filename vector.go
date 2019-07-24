package golangds

import "fmt"

// Vector data structure
type Vector struct {
	storage
}

// At func
func (v *Vector) At(index uint) (interface{}, error) {
	if index >= v.Size() || index < 0 {
		return nil, fmt.Errorf("Index '%d is out of limits", index)
	}
	if v.isConcurrent {
		v.mutex.RLock()
		defer v.mutex.RUnlock()
	}
	return *v.store[index], nil
}

// Front func
func (v *Vector) Front() (interface{}, error) {
	if v.Empty() {
		return nil, fmt.Errorf("The Vector is empty, it should have at least one item")
	}
	if v.isConcurrent {
		v.mutex.RLock()
		defer v.mutex.RUnlock()
	}
	return *v.store[len(v.store)-1], nil
}

// Back func
func (v *Vector) Back() (interface{}, error) {
	if v.Empty() {
		return nil, fmt.Errorf("The Vector is empty, it should have at least one item")
	}
	if v.isConcurrent {
		v.mutex.RLock()
		defer v.mutex.Unlock()
	}
	return *v.store[0], nil
}

// Assign func
func (v *Vector) Assign(n uint, value *interface{}) {
	if v.isConcurrent {
		v.mutex.Lock()
		defer v.mutex.Unlock()
	}
	for i := uint(0); i < n; i++ {
		v.store = append(v.store, value)
	}
}
