package main

import "fmt"

func main() {
	var number int
	var pointerOfInt *int
	var doublePointer **int

	//var decimal float64
	//decimal = 10.4

	number = 10
	pointerOfInt = &number
	doublePointer = &pointerOfInt

	fmt.Println("number:", number)
	fmt.Println("pointerOfInt:", pointerOfInt)
	fmt.Println("value pointed by pointerOfInt:", *pointerOfInt)
	fmt.Println("doublePointer:", doublePointer)
	fmt.Println("value pointed by doublePointer:", *doublePointer)
}
