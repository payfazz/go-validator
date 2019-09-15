package validator

import (
	"encoding/json"
	"github.com/payfazz/go-validator/validator/validation"

	en "github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

//Validate validator wrapper struct for validator.v9 Validate and universal-translator
type Validator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

//New create validator with default messages
func New() *Validator {
	val := validator.New()

	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

	v := &Validator{
		Validate: val,
		Trans:    trans,
	}

	v.RegisterMessages(map[string]string{
		"required": "{field} is required",
		"min":      "{field} min {param}",
		"max":      "{field} max {param}",
	})

	_ = v.Validate.RegisterValidation("date_rfc3339", validation.RFC3339)

	return v
}

//ValidateStruct validate struct
func (v *Validator) ValidateStruct(s interface{}) error {
	if nil == s {
		return nil
	}

	err := v.Validate.Struct(s)

	if nil == err {
		return nil
	}

	valerr, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}

	messages := valerr.Translate(v.Trans)
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
