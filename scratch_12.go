package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	items *list.List
}

func (q Queue) Enqueue(a any) {
	q.items.PushBack(a)
}

func (q Queue) Dequeue() any {
	if q.items.Len() == 0 {
		return -1
	}
	front := q.items.Front()
	val := q.items.Remove(front)
	return val
}

func main() {
	q := Queue{list.New()}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Dequeue(), q.Dequeue(), q.Dequeue())
	fmt.Println(q.Dequeue(), q.Dequeue(), q.Dequeue())
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.Dequeue(), q.Dequeue(), q.Dequeue())
}
