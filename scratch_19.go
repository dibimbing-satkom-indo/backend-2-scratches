package main

type Shape struct{
	a int
	b string
	c uint
}

type ShapeFactory struct{}

func (s ShapeFactory) createShape() Shape {
	return Shape{
		a: -100,
		b: "ok",
		c: 1000,
	}
}

func main() {
	factory := ShapeFactory{}
	factory.createShape()
}
