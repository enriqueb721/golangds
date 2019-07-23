package golangds

import "fmt"

// Queue data structure
type Queue struct {
	storage
}

// Front func
func (q *Queue) Front() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	if q.isConcurrent {
		q.mutex.RLock()
		defer q.mutex.RUnlock()
	}
	return *q.store[0], nil
}

// Back func
func (q *Queue) Back() (interface{}, error) {
	if q.Empty() {
		return nil, fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	if q.isConcurrent {
		q.mutex.RLock()
		defer q.mutex.RUnlock()
	}
	return *q.store[len(q.store)-1], nil
}

// Push func
func (q *Queue) Push(element interface{}) {
	if q.isConcurrent {
		q.mutex.Lock()
		defer q.mutex.Unlock()
	}
	q.store = append(q.store, &element)
}

// Pop func
func (q *Queue) Pop() error {
	if q.Empty() {
		return fmt.Errorf("The Queue is empty, it should have at least one item")
	}
	if q.isConcurrent {
		q.mutex.Lock()
		defer q.mutex.Unlock()
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
