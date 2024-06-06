package main

import (
	"fmt"
)

var pl = fmt.Println

type Queue struct {
	Values []int
}

func NewQueue() *Queue {
	return &Queue{
		Values: make([]int, 0),
	}
}

func (q *Queue) Enqueue(item int) {
	q.Values = append(q.Values, item)
}

func (q *Queue) Dequeue() int {
	if q.isEmpty() {
		pl("Queue is Empty!!")
		return 0
	}
	item := q.Values[0]
	q.Values = q.Values[1:]
	return item
}

func (q *Queue) isEmpty() bool {
	return len(q.Values) == 0
}

func main() {
	qu := NewQueue()
	qu.Dequeue()
	qu.Enqueue(1)
	qu.Enqueue(14)
	qu.Enqueue(12)
	pl(qu.Dequeue())
	pl(qu.Dequeue())
	pl(qu.Dequeue())
	pl(qu.Dequeue())

}
