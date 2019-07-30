package golangds

import (
	"fmt"

	"github.com/alexandrenriq/golangds/util"
)

// Deque data structure
type Deque struct {
	gdsutil.Storage
}

// At func
func (d *Deque) At(index uint) (interface{}, error) {
	if d.IsConcurrent() {
		defer d.RWMutex.RUnlock()
	}
	if index >= d.SizeWithoutRUnlock() || index < 0 {
		return nil, fmt.Errorf("Index '%d' is out of limits", index)
	}
	return *d.Store()[index], nil
}

// Front func
func (d *Deque) Front() (interface{}, error) {
	if d.IsConcurrent() {
		defer d.RWMutex.RUnlock()
	}
	if d.EmptyWithoutRUnlock() {
		return nil, fmt.Errorf("The Deque is empty, it should have at least one item")
	}
	return *d.Store()[len(d.Store())-1], nil
}

// Back func
func (d *Deque) Back() (interface{}, error) {
	if d.IsConcurrent() {
		defer d.RWMutex.RUnlock()
	}
	if d.EmptyWithoutRUnlock() {
		return nil, fmt.Errorf("The Deque is empty, it should have at least one item")
	}
	return *d.Store()[0], nil
}

// Assign func
func (d *Deque) Assign(n uint, value *interface{}) {
	if d.IsConcurrent() {
		d.RWMutex.Lock()
		defer d.RWMutex.Unlock()
	}
	for i := uint(0); i < n; i++ {
		d.SetStore(append(d.Store(), value))
	}
}

// PushBack func
func (d *Deque) PushBack(element interface{}) {
	if d.IsConcurrent() {
		d.RWMutex.Lock()
		defer d.RWMutex.Unlock()
	}
	d.SetStore(append(d.Store(), &element))
}

// PushFront func
func (d *Deque) PushFront(element interface{}) {
	if d.IsConcurrent() {
		d.RWMutex.Lock()
		defer d.RWMutex.Unlock()
	}
	var temp []*interface{}
	temp = append(temp, &element)
	d.SetStore(append(temp, d.Store()...))
}

// PopBack func
func (d *Deque) PopBack() error {
	if d.IsConcurrent() {
		defer d.RWMutex.Unlock()
	}
	if d.EmptyWithoutUnlock() {
		return fmt.Errorf("The Deque is empty, it should have at least one item")
	}
	d.SetStore(d.Store()[:1])
	return nil
}

// PopFront func
func (d *Deque) PopFront() error {
	if d.IsConcurrent() {
		defer d.RWMutex.Unlock()
	}
	if d.EmptyWithoutUnlock() {
		return fmt.Errorf("The Deque is empty, it should have at least one item")
	}
	d.SetStore(d.Store()[1:])
	return nil
}

// Insert func
func (d *Deque) Insert(index uint, element *interface{}) error {
	if d.IsConcurrent() {
		defer d.RWMutex.Unlock()
	}
	if index > d.SizeWithoutUnlock() {
		return fmt.Errorf("The index %d should not be greater than the current size %d", index, d.SizeWithoutConcurrency())
	}
	if index == d.SizeWithoutConcurrency() {
		d.SetStore(append(d.Store(), element))
		return nil
	}
	left := d.Store()[:index]
	right := d.Store()[index:]
	left = append(left, element)
	d.SetStore(append(left, right...))
	return nil
}

// Erase func
func (d *Deque) Erase(index uint) error {
	if d.IsConcurrent() {
		defer d.RWMutex.Unlock()
	}
	if index >= d.SizeWithoutUnlock() {
		return fmt.Errorf("The index %d should not be greater or equal than the current size %d", index, d.SizeWithoutConcurrency())
	}
	d.SetStore(append(d.Store()[:index], d.Store()[index+1:]...))
	return nil
}

// Swap func
func (d *Deque) Swap(anotherDeque *Deque) {
	temp := anotherDeque
	anotherDeque = d
	d = temp
}
