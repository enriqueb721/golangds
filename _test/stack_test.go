package test

import (
	"log"
	"testing"

	"github.com/alexandrenriq/golangds"
)

func TestStackPush(t *testing.T) {
	st := golangds.Stack{}
	st.Push(1)
	st.Push("asd")
	//for !st.Empty() {
	b, _ := st.Top()
	log.Print("hola ", b, " ", st.Size())
	st.Pop()
	b, _ = st.Top()
	log.Print("hola ", b, " ", st.Size())
	st.Pop()
	log.Fatal("rip ", st.Size(), " ", st.Empty())
	//}
}
