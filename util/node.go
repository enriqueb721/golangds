package gdsutil

// SimpleNode data structure
type SimpleNode struct {
	Mutex
	value interface{}
	next  *SimpleNode
}

// GetValue func
func (n *SimpleNode) GetValue() interface{} {
	return n.value
}

// GetNext func
func (n *SimpleNode) GetNext() *SimpleNode {
	return n.next
}

// SetValue func
func (n *SimpleNode) SetValue(value interface{}) {
	n.value = value
}

// SetNext func
func (n *SimpleNode) SetNext(next *SimpleNode) {
	n.next = next
}
