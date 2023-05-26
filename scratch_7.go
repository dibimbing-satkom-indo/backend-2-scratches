package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4}
	for _, number := range numbers {
		fmt.Println(number)
	}
}
