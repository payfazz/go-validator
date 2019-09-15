package main

import (
	"fmt"

	"github.com/payfazz/go-validator/validator"
)

//Product dummy struct for example
type Product struct {
	Name  string  `validate:"required,max=13"`
	Price float64 `validate:"required,min=0"`
}

func main() {
	validate := validator.New()

	customMessages := map[string]string{
		"required": "{field} must be filled",
		"min":      "{field} minimal {param}, your value is '{value}'",
		"max":      "{field} maximal {param}, your value is '{value}'",
	}
	validate.RegisterMessages(customMessages)

	product := &Product{}

	err := validate.ValidateStruct(product)
	fmt.Println(err)

	product.Name = "very long string for product name"
	product.Price = -1000

	err = validate.ValidateStruct(product)
	fmt.Println(err)
}
