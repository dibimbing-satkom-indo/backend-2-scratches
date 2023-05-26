package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func calculate(s Shape) {
	fmt.Println(s.Area())
}

type InvalidShape struct {}

func (is InvalidShape) Area() int  {
	return 0
}

type Rectangle struct {
	Side float64
}

func (r Rectangle) Area() float64 {
	return r.Side * r.Side
}

func main() {
	crc := Circle{100}
	//is := InvalidShape{}
	rect := Rectangle{100}

	calculate(crc)
	//calculate(is)
	calculate(rect)
}
