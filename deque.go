package golangds

import "fmt"

// Deque data structure
type Deque struct {
	storage
}

// At func
func (d *Deque) At(index uint) (interface{}, error) {
	if index >= d.Size() || index < 0 {
		return nil, fmt.Errorf("Index '%d' is out of limits", index)
	}
	return d.store[index], nil
}

// Front func
func (d *Deque) Front() (interface{}, error) {
	if d.Empty() {
		return nil, fmt.Errorf("The Deque is empty, it should have at least one item")
	}
	return d.store[len(d.store)-1], nil
}

// Back func
func (d *Deque) Back() (interface{}, error) {
	if d.Empty() {
		return nil, fmt.Errorf("The Deque is empty, it should have at least one item")
	}
	return d.store[0], nil
}

// Assign func
func (d *Deque) Assign(n uint, value *interface{}) {
	for i := uint(0); i < n; i++ {
		d.store = append(d.store, value)
	}
}

// PushBack func
func (d *Deque) PushBack(element interface{}) {
	d.store = append(d.store, &element)
}

// PushFront func
func (d *Deque) PushFront(element interface{}) {
	var temp []*interface{}
	temp = append(temp, &element)
	d.store = append(temp, d.store...)
}

// PopBack func
func (d *Deque) PopBack() error {
	if d.Empty() {
		return fmt.Errorf("The Deque is empty, it should have at least one item")
	}
	d.store = d.store[:1]
	return nil
}

// PopFront func
func (d *Deque) PopFront() error {
	if d.Empty() {
		return fmt.Errorf("The Deque is empty, it should have at least one item")
	}
	d.store = d.store[1:]
	return nil
}

// Insert func
func (d *Deque) Insert(index uint, element *interface{}) error {
	if index > d.Size() {
		return fmt.Errorf("The index %d should not be greater than the current size %d", index, d.Size())
	}
	if index == d.Size() {
		d.store = append(d.store, element)
		return nil
	}
	left := d.store[:index]
	right := d.store[index:]
	left = append(left, element)
	d.store = append(left, right...)
	return nil
}

// Erase func
func (d *Deque) Erase(index uint) error {
	if index >= d.Size() {
		return fmt.Errorf("The index %d should not be greater or equal than the current size %d", index, d.Size())
	}
	d.store = append(d.store[:index], d.store[index+1:]...)
	return nil
}

// Swap func
func (d *Deque) Swap(otherDeque *Deque) {
	temp := otherDeque
	otherDeque = d
	d = temp
}
