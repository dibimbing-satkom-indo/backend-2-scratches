package main

type DataStorage struct {
	validator DataValidator
}

func (ds DataStorage) Save(item int) {
	if ds.validator.Validate(item) {
		//
	}
}

type DataValidator interface {
	Validate(item int) bool
}

type PositiveNegativeValidator struct {}

func (v PositiveNegativeValidator) Validate(item int) bool {
	if item > 0 {
		return true
	}

	return false
}

func (v PositiveNegativeValidator) Len(item any) int {
	return 0
}

type AlwaysTrueDataValidator struct{}

func (dv AlwaysTrueDataValidator) Validate(item int) bool {
	return true
}

func main() {
	ds := DataStorage{validator: PositiveNegativeValidator{}}
	ds.Save(100)
}
