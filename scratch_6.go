package main

type controller struct {
	model string
}

func (c controller) create() {

}

func (c controller) read() {

}

func main() {
	ctrl := controller{
		model: "model",
	}
	ctrl.create()
	ctrl.read()
}
