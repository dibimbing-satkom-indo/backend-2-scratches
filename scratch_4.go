package main

import "fmt"

func main() {
	num := 42
	var ptr *int

	ptr = &num

	fmt.Println("value of num:", num)
	fmt.Println("address of num:", ptr)

	var num2 int

	*ptr = 12

	ptr = &num2

	num2 = 100
	*ptr = 20

	fmt.Println("addressed pointer value:", *ptr)
	fmt.Println("num value:", num)
	fmt.Println("num2 value:", num2)
}
