# go-validator

Golang struct tag validator based on [https://github.com/go-playground/validator](https://github.com/go-playground/validator).

Struct with **validate** tag:
```
type Product struct {
	Name  string  `validate:"required,max=13"`
	Price float64 `validate:"required,min=0"`
}
```

Executing validator:
```
validate := validator.New()

product := &Product{}

err := validate.ValidateStruct(product)
fmt.Println(err)
```

Override global default tag-level messages:
```
validate := validator.New()

customMessages := map[string]string{
	"required": "{field} must be filled",
	"min":      "{field} minimal {param}, your value is '{value}'",
	"max":      "{field} maximal {param}, your value is '{value}'",
}
validate.RegisterMessages(customMessages)
```

Override for spesific validation execution field-level or tag-level messages with decorator:
```
validate := validator.New()

customMessages := map[string]string{
	"Name.required": "{field} must be filled",
}

product := &Product{}

err := validate.WithCustomFieldMessages(customMessages).ValidateStruct(product)
 ```
