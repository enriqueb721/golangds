package gdselem

import (
	"fmt"

	"github.com/alexandrenriq/golangds/util"
)

// Deque data structure
type Deque struct {
	gdsutil.Storage
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

// Swap func
func (d *Deque) Swap(anotherDeque *Deque) {
	temp := anotherDeque
	anotherDeque = d
	d = temp
}
