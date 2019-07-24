package test

import (
	"testing"
	"time"

	"github.com/alexandrenriq/golangds"
)

func TestPush(t *testing.T) {
	st := golangds.Stack{}
	for i := 0; i < 10000000; i++ {
		st.Push(i)
	}
	if st.Size() != 10000000 {
		t.Error("The stack should have 10000000 values")
	}
	for i := 10000000 - 1; i >= 0; i-- {
		val, err := st.Top()
		if err != nil {
			t.Error("An error has been caught in Top function")
		}
		if val != i {
			t.Errorf("The current top value of the stack should be %d instead of %d", i, val)
		}
		err = st.Pop()
		if err != nil {
			t.Error("An error has been caught in Pop function")
		}
	}
}

func TestTopEmpty(t *testing.T) {
	st := golangds.Stack{}
	_, err := st.Top()
	if err != nil {
		t.Logf("The error has been caught, %e", err)
		return
	}
	t.Error("There is no caught of error")
}

func TestEmpty1(t *testing.T) {
	st := golangds.Stack{}
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	if st.Empty() {
		t.Error("The stack should not be empty")
	}
	err := st.Pop() // 5
	if st.Empty() {
		t.Error("The stack should not be empty")
	}
	if err != nil {
		t.Error("An error has been caught in Pop function")
	}
	err = st.Pop() // 4
	if st.Empty() {
		t.Error("The stack should not be empty")
	}
	if err != nil {
		t.Error("An error has been caught in Pop function")
	}
	err = st.Pop() // 3
	if st.Empty() {
		t.Error("The stack should not be empty")
	}
	if err != nil {
		t.Error("An error has been caught in Pop function")
	}
	err = st.Pop() // 2
	if st.Empty() {
		t.Error("The stack should not be empty")
	}
	if err != nil {
		t.Error("An error has been caught in Pop function")
	}
	err = st.Pop() // 1
	if !st.Empty() {
		t.Error("The stack should be empty")
	}
	if err != nil {
		t.Error("An error has been caught in Pop function")
	}
}

func TestEmpty2(t *testing.T) {
	st := golangds.Stack{}
	for i := 0; i < 10000000; i++ {
		st.Push(10000000)
	}
	for i := 0; i < 10000000; i++ {
		err := st.Pop()
		if err != nil {
			t.Error("An error has been caught in Pop function")
		}
		if i < 10000000-1 && st.Empty() {
			t.Error("The stack should not be empty")
		}
	}
	if !st.Empty() {
		t.Error("The stack should be empty")
	}
}

func TestTopNotEmpty(t *testing.T) {
	st := golangds.Stack{}
	for i := 0; i < 10000000; i++ {
		st.Push(10000000)
	}
	for !st.Empty() {
		_, err := st.Top()
		if err != nil {
			t.Error("There is an error in Top function")
		}
		err = st.Pop()
		if err != nil {
			t.Error("An error has been caught in Pop function")
		}
	}
}

func TestClearThenPush(t *testing.T) {
	st := golangds.Stack{}
	for i := 0; i < 10; i++ {
		st.Push(i)
	}
	st.Clear()
	if st.Size() > 0 {
		t.Error("The stack should be empty")
	}
	for i := 0; i < 10; i++ {
		st.Push(i)
	}
	if st.Size() != 10 {
		t.Error("The stack should have 10 values")
	}
}

// TODO: test swap function

func TestConcurrentPushAndPop(t *testing.T) {
	st := golangds.Stack{}
	st.EnableConcurrency()
	endF1 := false
	endF2 := false
	go func() {
		for !st.Empty() || !endF2 {
			if !st.Empty() {
				t.Log("Empty")
				val, err := st.Top()
				if err != nil {
					t.Error("There is an error in Top function")
				}
				err = st.Pop()
				if err != nil {
					t.Error("There is an error in Pop function")
				}
				t.Logf("Empty val: %d", val)
			}
		}
		endF1 = true
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			st.Push(i)
			t.Log(i)
		}
		endF2 = true
	}()
	time.Sleep(2 * time.Second)
	if !endF1 || !endF2 {
		t.Error("The concurrent functions did not finish")
		t.Log(endF1, endF2)
	}
}
