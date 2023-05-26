package main

import "fmt"

func printOddNumber(ch chan int) {
	for i:=1; i < 10; i=i+2 {
		fmt.Println("   ",i)
	}
	ch <- 1
}

func printEvenNumber(ch chan int) {
	for i:=2; i < 10; i=i+2 {
		fmt.Println(i)
	}
	//ch <- 2
}

func printNumber(ch chan int, id int) {
	for i:=1; i < 10; i=i+2 {
		fmt.Println("   ",i)
	}
	ch <- id
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	go printOddNumber(ch)
	go printOddNumber(ch2)
	//printEvenNumber(ch)
	fmt.Println("Finished:", <-ch)
	fmt.Println("Finished:", <-ch2)
	//fmt.Println("Finished:", <-ch)
	//fmt.Println("Finished:", <-ch)
	//
	//for i:=1; i < 11; i++ {
	//	go printNumber(ch, i)
	//}
	//
	//for i:=1; i < 11; i++ {
	//	fmt.Println("Finished:", <-ch)
	//}
}
