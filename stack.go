package golangds

import "fmt"

// Stack data structure
type Stack struct {
	storage
}

// Top func
func (s *Stack) Top() (interface{}, error) {
	if s.mutex.isConcurrent {
		defer s.mutex.RUnlock()
	}
	if s.emptyWithoutRUnlock() { // Mantain the RLock function until the end of the current function
		return nil, fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	return *s.store[len(s.store)-1], nil
}

// Bottom func
func (s *Stack) Bottom() (interface{}, error) {
	if s.mutex.isConcurrent {
		defer s.mutex.RUnlock()
	}
	if s.emptyWithoutRUnlock() {
		return nil, fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	return *s.store[0], nil
}

// Push func
func (s *Stack) Push(element interface{}) {
	if s.mutex.isConcurrent {
		s.mutex.Lock()
		defer s.mutex.Unlock()
	}
	s.store = append(s.store, &element)
}

// Pop func
func (s *Stack) Pop() error {
	if s.mutex.isConcurrent {
		defer s.mutex.Unlock()
	}
	if s.emptyWithoutUnlock() {
		return fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	s.store = s.store[:len(s.store)-1]
	return nil
}

// Swap func
func (s *Stack) Swap(anotherStack *Stack) {
	temp := anotherStack
	anotherStack = s
	s = temp
}
