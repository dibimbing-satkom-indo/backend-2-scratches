package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}

	fmt.Println()
}

func main() {
	l := list.New()
	l.PushFront(1)
	l.PushFront("str")
	l.PushFront(nil)
	l.PushFront(struct {
		name string
	}{
		name: "str",
	})

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println(l.Back().Value)

	ll := &LinkedList{}

	ll.Append(10)
	ll.Append(20)

	ll.Display()
}
