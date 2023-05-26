package main

import (
	"errors"
	"fmt"
	"log"
)

func divideByTwo(val int) (*int, error) {
	if val > 0 {
		val = val/2
		return &val, nil
	}

	return &val, errors.New("won't divide zero")
}

func main() {
	result, err := divideByTwo(10)
	if err == nil {
		fmt.Println(*result)
	} else {
		log.Println(err)
	}

	result2, _ := divideByTwo(210)
	log.Println(*result2)

	result3, err := divideByTwo(0)
	if err == nil {
		fmt.Println(*result3)
	} else {
		log.Println(err)
	}
}
