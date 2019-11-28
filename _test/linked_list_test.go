package test

import (
	"testing"

	"github.com/alexandrenriq/golangds/list"
)

func TestLinkedList_Insert(t *testing.T) {
	l := gdslist.LinkedList{}
	for i := 0; i < 1e3; i++ {
		t.Logf("i: %d", i)
		l.Insert(i)
	}

}
