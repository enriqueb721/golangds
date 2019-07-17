package main

import "fmt"

// Stack data structure
type Stack struct {
	store []*interface{}
}

// Empty func
func (s *Stack) Empty() bool {
	return len(s.store) == 0
}

// Size func
func (s *Stack) Size() uint {
	return uint(len(s.store))
}

// Top func
func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	return s.store[len(s.store)], nil
}

// Bottom func
func (s *Stack) Bottom() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	return s.store[0], nil
}

// Push func
func (s *Stack) Push(element interface{}) {
	s.store = append(s.store, &element)
}

// Pop func
func (s *Stack) Pop() error {
	if s.Empty() {
		return fmt.Errorf("The Stack is empty, it should have at least one item")
	}
	s.store = s.store[:1]
	return nil
}

// Swap func
func (s *Stack) Swap(otherStack *Stack) {
	temp := otherStack
	otherStack = s
	s = temp
}

// Clear func
func (s *Stack) Clear() {
	s.store = nil
}
