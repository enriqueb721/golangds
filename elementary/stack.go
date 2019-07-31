package gdselem

import (
	"fmt"

	"github.com/alexandrenriq/golangds/util"
)

// Stack data structure
type Stack struct {
	gdsutil.Storage
}

// Top func
func (s *Stack) Top() (interface{}, error) {
	if s.IsConcurrent() {
		defer s.RWMutex.RUnlock()
	}
	if s.EmptyWithoutRUnlock() { // Mantain the RLock function until the end of the current function
		return nil, fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	return *s.Store()[len(s.Store())-1], nil
}

// Bottom func
func (s *Stack) Bottom() (interface{}, error) {
	if s.IsConcurrent() {
		defer s.RWMutex.RUnlock()
	}
	if s.EmptyWithoutRUnlock() {
		return nil, fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	return *s.Store()[0], nil
}

// Push func
func (s *Stack) Push(element interface{}) {
	if s.IsConcurrent() {
		s.RWMutex.Lock()
		defer s.RWMutex.Unlock()
	}
	s.SetStore(append(s.Store(), &element))
}

// Pop func
func (s *Stack) Pop() error {
	if s.IsConcurrent() {
		defer s.RWMutex.Unlock()
	}
	if s.EmptyWithoutUnlock() {
		return fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	s.SetStore(s.Store()[:len(s.Store())-1])
	return nil
}

// Swap func
func (s *Stack) Swap(anotherStack *Stack) {
	temp := anotherStack
	anotherStack = s
	s = temp
}
