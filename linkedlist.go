package main

import "fmt"

// Node represents a node in the linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList is a generic linked list
type LinkedList[T any] struct {
	Head *Node[T]
}

// AddNode adds a new node with the given value to the linked list
func (ll *LinkedList[T]) AddNode(value T) {
	newNode := &Node[T]{Value: value, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// PrintList prints the elements of the linked list
func (ll *LinkedList[T]) PrintList() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	// Creating a linked list of integers
	intList := LinkedList[int]{}
	intList.AddNode(1)
	intList.AddNode(2)
	intList.AddNode(3)

	// Printing the integer linked list
	fmt.Println("Integer Linked List:")
	intList.PrintList()

	// Creating a linked list of strings
	strList := LinkedList[string]{}
	strList.AddNode("a")
	strList.AddNode("b")
	strList.AddNode("c")

	// Printing the string linked list
	fmt.Println("\nString Linked List:")
	strList.PrintList()
}
