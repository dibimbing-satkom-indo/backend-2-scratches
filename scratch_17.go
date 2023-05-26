package main

type Person struct {
	name string
	telephoneNumber TelephoneNumber
}

func (p Person) GetTelephoneNumber() string {
	return p.telephoneNumber.GetTelephoneNumber()
}

type TelephoneNumber struct {
	officeAreaNumber string
	officeNumber string
}

func (t TelephoneNumber) GetTelephoneNumber() string {
	return t.officeAreaNumber + "-" +t.officeNumber
}

func main() {
	
}
