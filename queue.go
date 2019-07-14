package main

import "fmt"

// Queue data structure
type Queue struct {
	store []*interface{}
}

// Empty func
func (q *Queue) Empty() bool {
	return len(q.store) == 0
}

// Size func
func (q *Queue) Size() uint {
	return uint(len(q.store))
}

// Front func
func (q *Queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	return q.store[0], nil
}

// Back func
func (q *Queue) Back() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	return q.store[len(q.store)-1], nil
}

// Push func
func (q *Queue) Push(element interface{}) {
	q.store = append(q.store, &element)
}

// Pop func
func (q *Queue) Pop() error {
	if q.Empty() {
		return fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	q.store = q.store[1:]
	return nil
}

// Swap func
func (q *Queue) Swap(otherQueue *Queue) {
	temp := otherQueue
	otherQueue = q
	q = temp
}

// TODO: Clear function