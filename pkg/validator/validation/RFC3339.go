package validation

import (
	"fmt"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

//RFC3339 validation
func RFC3339(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())

	if nil != err {
		fmt.Println(err)
		return false
	}

	return true
}
