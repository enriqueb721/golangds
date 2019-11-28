package test

import (
	"testing"

	"github.com/alexandrenriq/golangds/elementary"
)

func TestQueue_Push(t *testing.T) {
	q := gdselem.Queue{}
	for i := 0; i < 1e7; i++ {
		q.Push(i)
	}
	if q.Size() != 1e7 {
		t.Error("The queue should have 1e7 values")
	}
	for i := 0; i < 1e7; i++ {
		val, err := q.Front()
		if err != nil {
			t.Error("An error has been caught in Front function")
		}
		if val != i {
			t.Errorf("The current top value of the stack should be %d instead of %d", i, val)
		}
		err = q.Pop()
		if err != nil {
			t.Error("An error has been caught in Pop function")
		}
	}
}
