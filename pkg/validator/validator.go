package validator

import (
	"fmt"

	en "github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validate *validator.Validate
	trans    ut.Translator
}

func New() *Validator {
	validate := validator.New()

	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

	RegisterTranslation(validate, trans, "required", "{field} is required")

	return &Validator{
		validate: validate,
		trans:    trans,
	}
}

type validatorError struct {
	errors []string
}

func (e validatorError) Error() string {
	if len(e.errors) == 1 {
		return e.errors[0]
	}

	result := ""
	for _, err := range e.errors {
		if result != "" {
			result += ","
		}
		result += fmt.Sprintf(`"%s"`, err)
	}
	return result
}

func (v *Validator) ValidateStruct(s interface{}) error {
	err := v.validate.Struct(s)

	if err == nil {
		return nil
	}

	messages := err.(validator.ValidationErrors).Translate(v.trans)

	var errors []string
	for _, message := range messages {
		errors = append(errors, message)
	}

	return &validatorError{
		errors: errors,
	}
}
