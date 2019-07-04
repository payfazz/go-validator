package validator

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

//CustomFieldMessages key-value pair for tag and translation
//Tag format: [struct field name].[validation tag]
type CustomFieldMessages map[string]string

//FieldCustom wrap validator and custom field messages
type FieldCustom struct {
	Val      *Validator
	Messages CustomFieldMessages
}

//WithCustomFieldMessages return FieldCustom wrapper
func (val *Validator) WithCustomFieldMessages(messages CustomFieldMessages) *FieldCustom {
	return &FieldCustom{
		Val:      val,
		Messages: messages,
	}
}

//ValidateStruct validate struct and translate
func (val *FieldCustom) ValidateStruct(obj interface{}) error {
	verrs := val.Val.Validate.Struct(obj).(validator.ValidationErrors)

	messages := make(map[string]string)
	for _, verr := range verrs {
		key := verr.Field() + "." + verr.Tag()
		if translation, ok := val.Messages[key]; ok {
			messages[verr.StructNamespace()] = translate(verr, translation)
		} else {
			messages[verr.StructNamespace()] = verr.Translate(val.Val.Trans)
		}
	}

	return &validatorError{
		messages: messages,
	}
}

func translate(fe validator.FieldError, format string) string {
	tags, translation := getAndReplaceTranslationKeywordsSprintf(format)

	tagsVal := getParamByTags(tags, fe)

	var params []interface{}
	for _, val := range tagsVal {
		params = append(params, val)
	}

	result := fmt.Sprintf(translation, params...)
	return result
}
