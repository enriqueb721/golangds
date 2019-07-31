package gdselem

import (
	"fmt"

	"github.com/alexandrenriq/golangds/util"
)

// Queue data structure
type Queue struct {
	gdsutil.Storage
}

// Front func
func (q *Queue) Front() (interface{}, error) {
	if q.IsConcurrent() {
		defer q.RWMutex.RUnlock()
	}
	if q.EmptyWithoutRUnlock() {
		return nil, fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	return *q.Store()[0], nil
}

// Back func
func (q *Queue) Back() (interface{}, error) {
	if q.IsConcurrent() {
		defer q.RWMutex.RUnlock()
	}
	if q.EmptyWithoutRUnlock() {
		return nil, fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	return *q.Store()[len(q.Store())-1], nil
}

// Push func
func (q *Queue) Push(element interface{}) {
	if q.IsConcurrent() {
		q.RWMutex.Lock()
		defer q.RWMutex.Unlock()
	}
	q.SetStore(append(q.Store(), &element))
}

// Pop func
func (q *Queue) Pop() error {
	if q.IsConcurrent() {
		defer q.RWMutex.Unlock()
	}
	if q.EmptyWithoutUnlock() {
		return fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	q.SetStore(q.Store()[1:])
	return nil
}

// Swap func
func (q *Queue) Swap(anotherQueue *Queue) {
	temp := anotherQueue
	anotherQueue = q
	q = temp
}
