package test

import (
	"testing"

	"github.com/alexandrenriq/golangds"
)

func hola(t *testing.T) {
	t.Log("hola :'v")

}

func TestQueue(t *testing.T) {
	a := golangds.Stack{}
	a.Clear()
	t.Log("hola :'v")
	hola(t)
}
