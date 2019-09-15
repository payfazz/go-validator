package validator_test

import (
	"encoding/json"
	"github.com/payfazz/go-validator/validator"
	"testing"
)

type TestTranslateFieldStruct struct {
	A       string  `validate:"min=5"`
	B       string  `validate:"iscolor"`
	C       string  `validate:"min=3"`
	D       string  `validate:"min=13"`
	Float32 float32 `validate:"min=100"`
	Int     int     `validate:"min=100"`
}

func TestTranslateField(t *testing.T) {
	obj := &TestTranslateFieldStruct{}

	custom := map[string]string{
		"A.min":       "{field} minimal {param}!",
		"C.min":       "{field} length please at least {param}!",
		"Float32.min": "your value {value} is invalid!",
		"Int.min":     "your value {value} is invalid!",
	}

	val := validator.New()

	err := val.WithCustomFieldMessages(custom).ValidateStruct(obj)

	var data map[string]string
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestTranslateFieldStruct.A"] != "A minimal 5!" {
		t.Log(err)
		t.Error("wrong custom field message")
	}

	err = val.ValidateStruct(obj)
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestTranslateFieldStruct.A"] != "A min 5" {
		t.Log(err)
		t.Error("wrong default field message")
	}

	custom = map[string]string{
		"min": "you must input {field} with minimal {param}",
	}

	err = val.WithCustomFieldMessages(custom).ValidateStruct(obj)
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestTranslateFieldStruct.A"] == "A min 5" {
		t.Log(err)
		t.Error("wrong custom field message")
	}
}
