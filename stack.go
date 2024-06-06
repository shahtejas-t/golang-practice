package main

import (
	"fmt"
)

var pl = fmt.Println

type Stack struct {
	values []int
	top    int
}

func NewStack() *Stack {
	return &Stack{
		values: make([]int, 0),
		top:    -1,
	}
}

func (st *Stack) push(n int) {
	// pl("PUSH TOP :", st.top)
	st.top++
	st.values = append(st.values, n)
}

func (st *Stack) pop() int {
	if st.isEmpty() {
		pl("Stack is empty!!")
		return 0
	}
	// pl("POP TOP :", st.top)
	value := st.values[st.top]
	st.top--
	return value
}

func (st *Stack) isEmpty() bool {
	if st.top == -1 {
		return true
	}
	return false
}

func main() {
	var st = NewStack()
	pl(st.pop())
	st.push(1)
	st.push(17)
	st.push(4)
	st.push(9)
	st.push(3)
	pl(st.pop())
	pl(st.pop())
	pl(st.pop())
	pl(st.pop())
	pl(st.pop())
	pl(st.pop())
}
