package main

import (
	"fmt"
	"strings"
)

func FindIntersection(strArr []string) string {

	first := strings.Split(strArr[0], ", ")
	second := strings.Split(strArr[1], ", ")
	intersection := make([]string, 0)
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(second); j++ {
			if first[i] == second[j] {
				intersection = append(intersection, first[i])
				break
			}
		}
	}
	return strings.Join(intersection, ",")
}

func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(FindIntersection(readline()))

}