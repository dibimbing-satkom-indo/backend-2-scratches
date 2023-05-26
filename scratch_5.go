package main

import "fmt"

type incrementorFunc func(int) int

func applyOperation(name string, val int, operation incrementorFunc) int {
	fmt.Println("operation:", name)
	return operation(val)
}

func incrementByOne(val int) int {
	return val+1
}

func incrementByTwo(val int) int {
	return val+2
}

func main() {
	result := applyOperation("increment by one", 10, incrementByOne)

	var opt incrementorFunc

	opt = incrementByTwo
	result2 := applyOperation("increment by two", 20, opt)

	fmt.Println("result:", result)
	fmt.Println("result2:", result2)

	var map1 map[string]int
	map1 = make(map[string]int, 0)
	map2 := map[string]int {
		"satu": 1,
	}
	map3 := map[string]int {}

	map1["satu"] = 1
	map3["tiga"] = 3

	var ptr *int

	fmt.Println(map1, ptr, map1 == nil)
	fmt.Println("map1:", map1)
	fmt.Println("map2:", map2)
	fmt.Println("map3:", map3)
}
