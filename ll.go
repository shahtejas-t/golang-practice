package main

import (
	"fmt"
)

var pl = fmt.Println

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type Linklist[T any] struct {
	Head *Node[T]
}

func (ll *Linklist[T]) add(value T) {
	newNode := &Node[T]{Value: value, Next: nil}
	if ll.isEmpty() {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *Linklist[T]) printll() {
	if ll.isEmpty() {
		pl("List is empty")
		return
	}
	current := ll.Head
	for current != nil {
		pl(current.Value)
		current = current.Next
	}
	pl("nil")
}

func (ll *Linklist[T]) isEmpty() bool {
	return ll.Head == nil
}

func main() {
	ll := Linklist[int]{}
	ll.printll()
	ll.add(1)
	ll.add(5)
	ll.add(8)
	ll.add(44)
	ll.add(98)
	ll.printll()
}
