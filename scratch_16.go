package main

type Order struct {
	customer Customer
}

type Customer struct {
	name string
	address string
}

type RegisterFeature struct {}

func (f RegisterFeature) Register(cust Customer) {
	// register customer
}

func main() {
	customer := Customer{}
	feat := RegisterFeature{}
	feat.Register(customer)
}
