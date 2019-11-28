package gdslist

import (
	"github.com/alexandrenriq/golangds/util"
)

// LinkedList Data structure
type LinkedList struct {
	head *gdsutil.SimpleNode
	tail *gdsutil.SimpleNode
}

// Insert func
func (l *LinkedList) Insert(value interface{}) error {
	newNode := &gdsutil.SimpleNode{}
	newNode.SetValue(value)
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		newNode = nil
	} else {
		l.tail.SetNext(newNode)
		l.tail = newNode
	}
	return nil
}

// InsertToHead func
func (l *LinkedList) InsertToHead(value interface{}) error {
	newNode := &gdsutil.SimpleNode{}
	newNode.SetValue(value)
	newNode.SetNext(l.head)
	l.head = newNode
	return nil
}

// GetHead func
func (l *LinkedList) GetHead() (*gdsutil.SimpleNode, error) {
	return l.head, nil
}

// GetTail func
func (l *LinkedList) GetTail() (*gdsutil.SimpleNode, error) {
	return l.tail, nil
}
