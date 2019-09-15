package main

import (
	"fmt"

	"github.com/payfazz/go-validator/validator"
)

//Product dummy struct for example
type Product struct {
	Name string `validate:"required"`
	Type string `validate:"required"`
}

func main() {
	validate := validator.New()

	customMessages := map[string]string{
		"Name.required": "{field} must be filled",
	}

	product := &Product{}

	err := validate.WithCustomFieldMessages(customMessages).ValidateStruct(product)
	fmt.Println(err)

	err = validate.ValidateStruct(product)
	fmt.Println(err)
}
