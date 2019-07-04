package validator_test

import (
	"encoding/json"
	"testing"

	"github.com/payfazz/go-validator/pkg/validator"
)

type TestTranslationStruct struct {
	Number int `validate:"min=5,max=12"`
}

func TestTranslation(t *testing.T) {
	val := validator.New()
	val.RegisterTranslation(map[string]string{
		"min": "{field} is not valid, min: {param}",
		"max": "{field} is not valid, max: {param}",
	})

	obj := &TestTranslationStruct{
		Number: 123,
	}

	err := val.ValidateStruct(obj)

	var data map[string]string
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestTranslationStruct.Number"] != "Number is not valid, max: 12" {
		t.Log(data)
		t.Error("validation translation failed")
	}
}

func TestTranslationPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("must panic")
		}
	}()

	val := validator.New()
	val.RegisterTranslation(map[string]string{
		"abc": "{1} is not valid, min: {3}",
	})
}

func TestTranslationPanic2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("must panic")
		}
	}()

	val := validator.New()
	val.RegisterTranslation(map[string]string{
		"abc": "{1} is not valid, min: {3}",
	})
}

type TestTagStruct struct {
	A string `validate:"min=5"`
	B string `validate:"iscolor"`
}

func TestTag(t *testing.T) {
	obj := &TestTagStruct{}

	val := validator.New()
	val.RegisterTranslation(map[string]string{
		"min":     "{tag} {actualTag} {namespace} {structNamespace} {field} {structField} {value} {param}",
		"iscolor": "{tag} {actualTag} {namespace} {structNamespace} {field} {structField} {value} {param}",
	})

	err := val.ValidateStruct(obj)

	var data map[string]string
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestTagStruct.A"] != "min min TestTagStruct.A TestTagStruct.A A A  5" {
		t.Log(err)
		t.Error("wrong translation on min")
	}

	if data["TestTagStruct.B"] != "iscolor hexcolor|rgb|rgba|hsl|hsla TestTagStruct.B TestTagStruct.B B B  " {
		t.Log(err)
		t.Error("wrong translation on iscolor")
	}
}

func TestTranslateField(t *testing.T) {
	obj := &TestTagStruct{}

	custom := map[string]string{
		"A.min": "{field} minimal {param}!",
	}

	val := validator.New()

	err := val.WithCustomFieldMessage(custom).ValidateStruct(obj)

	var data map[string]string
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestTagStruct.A"] != "A minimal 5!" {
		t.Log(err)
		t.Error("wrong custom field translation")
	}
}
