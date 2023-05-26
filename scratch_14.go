package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data string
	next *Node
}

func (ll *LinkedList) Append(data string) {
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

func incrementByOne(val int) {
	val = val+1
}

func decrementByOne(val *int) {
	*val = *val-1
}

func multiplyByTwo(val int) int {
	return val * 2
}

//func (ll *LinkedList) ChangeSomething() {
//	ll.head = &Node{data: "node something", next: nil}
//}

func main() {
	ll := LinkedList{head: nil}
	ll.Append("satu")
	ll.Append("dua")
	ll.Append("tiga")
	fmt.Println(ll.head.data, ll.head.next.data, ll.head.next.next.data, ll.head.next.next.next)
	//ll.ChangeSomething()
	//fmt.Println(ll.head)
	//number := 10
	//incrementByOne(number)
	//fmt.Println(number)
	//decrementByOne(&number)
	//fmt.Println(number)
	//number = multiplyByTwo(number)
	//fmt.Println(number)
}
