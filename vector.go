package golangds

import "fmt"

// Vector data structure
type Vector struct {
	array
}

// Data func
func (v *Vector) Data() []*interface{} {
	return v.Store
}

// At func
func (v *Vector) At(index uint) (interface{}, error) {
	if v.mutex.isConcurrent {
		defer v.mutex.RUnlock()
	}
	if index >= v.sizeWithoutRUnlock() || index < 0 {
		return nil, fmt.Errorf("Index '%d is out of limits", index)
	}
	return *v.Store[index], nil
}

// Front func
func (v *Vector) Front() (interface{}, error) {
	if v.mutex.isConcurrent {
		defer v.mutex.RUnlock()
	}
	if v.emptyWithoutRUnlock() {
		return nil, fmt.Errorf("The Vector is empty, it should have at least one item")
	}
	return *v.Store[len(v.Store)-1], nil
}

// Back func
func (v *Vector) Back() (interface{}, error) {
	if v.mutex.isConcurrent {
		defer v.mutex.RUnlock()
	}
	if v.emptyWithoutRUnlock() {
		return nil, fmt.Errorf("The Vector is empty, it should have at least one item")
	}
	return *v.Store[0], nil
}

// Assign func
func (v *Vector) Assign(n uint, value *interface{}) {
	if v.mutex.isConcurrent {
		v.mutex.Lock()
		defer v.mutex.Unlock()
	}
	for i := uint(0); i < n; i++ {
		v.Store = append(v.Store, value)
	}
}

// Append func
func (v *Vector) Append(element interface{}) {
	if v.mutex.isConcurrent {
		v.mutex.Lock()
		defer v.mutex.Unlock()
	}
	v.Store = append(v.Store, &element)
}

// PopBack func
func (v *Vector) PopBack() error {
	if v.mutex.isConcurrent {
		defer v.mutex.Unlock()
	}
	if v.emptyWithoutUnlock() {
		return fmt.Errorf("The Vector is empty, it should have at least one item")
	}
	v.Store = v.Store[:1]
	return nil
}

// Insert func
func (v *Vector) Insert(index uint, element *interface{}) error {
	if v.mutex.isConcurrent {
		defer v.mutex.Unlock()
	}
	if index > v.sizeWithoutUnlock() {
		return fmt.Errorf("The index %d should not be greater than the current size %d", index, v.sizeWithoutConcurrency())
	}
	if index == v.sizeWithoutConcurrency() {
		v.Store = append(v.Store, element)
		return nil
	}
	left := v.Store[:index]
	right := v.Store[index:]
	left = append(left, element)
	v.Store = append(left, right...)
	return nil
}

// Erase func
func (v *Vector) Erase(index uint) error {
	if v.mutex.isConcurrent {
		defer v.mutex.Unlock()
	}
	if index >= v.sizeWithoutUnlock() {
		return fmt.Errorf("The index %d should not be greater or equal than the current size %d", index, v.sizeWithoutConcurrency())
	}
	v.Store = append(v.Store[:index], v.Store[index+1:]...)
	return nil
}

// Swap func
func (v *Vector) Swap(anotherVector *Vector) {
	temp := anotherVector
	anotherVector = v
	v = temp
}
