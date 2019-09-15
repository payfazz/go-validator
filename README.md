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
val := validator.New()

product := &Product{}

err := val.ValidateStruct(product)
fmt.Println(err)
```

Override global default tag-level messages:
```
val := validator.New()

customMessages := map[string]string{
	"required": "{field} must be filled",
	"min":      "{field} minimal {param}, your value is '{value}'",
	"max":      "{field} maximal {param}, your value is '{value}'",
}
val.RegisterMessages(customMessages)
```

Override for spesific validation execution field-level or tag-level messages with decorator:
```
val := validator.New()

customMessages := map[string]string{
	"Name.required": "{field} must be filled",
}

product := &Product{}

err := val.WithCustomFieldMessages(customMessages).ValidateStruct(product)
 ```

Validator v9 methods still can be use from Validate object.
```
import (
	"github.com/payfazz/go-validator/validator"
    validator_v9 "gopkg.in/go-playground/validator.v9"
)

type Test struct {
    Image string `validate:"type=jpg|type=png"`
}

func main() {
    val := validator.New()

    _ = val.Validate.RegisterValidation("type", func(f validator_v9.FieldLevel) bool {
        return strings.HasSuffix(f.Field().String(), f.Param())
    })
    
    err := val.ValidateStruct(Test{
        Image: "test.xyz",
    })
    
    fmt.Println(err)
}
```
