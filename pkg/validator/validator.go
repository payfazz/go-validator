package validator

import (
	"encoding/json"

	en "github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

//Validator validator wrapper struct for validator.v9 Validate and universal-translator
type Validator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

//New create validator with default translation
func New() *Validator {
	validate := validator.New()

	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

	v := &Validator{
		Validate: validate,
		Trans:    trans,
	}

	v.RegisterTranslation(map[string]string{
		"required": "{field} is required",
		"min":      "{field} must be minimal {param}",
		"max":      "{field} must be maximal {param}",
	})

	return v
}

//ValidateStruct validate struct and translate
func (v *Validator) ValidateStruct(s interface{}) error {
	err := v.Validate.Struct(s)

	if err == nil {
		return nil
	}

	messages := err.(validator.ValidationErrors).Translate(v.Trans)
	return &validatorError{
		messages: messages,
	}
}

type validatorError struct {
	messages validator.ValidationErrorsTranslations
}

func (e validatorError) Error() string {
	s, _ := json.Marshal(e.messages)
	return string(s)
}
