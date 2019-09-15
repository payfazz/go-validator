package validator_test

import (
	"encoding/json"
	"github.com/payfazz/go-validator/validator"
	"testing"

	"github.com/google/uuid"
)

type TestValidateRequiredStruct struct {
	Foo string    `validate:"required"`
	Bar uuid.UUID `validate:"required"`
}

func TestValidateRequired(t *testing.T) {
	val := validator.New()

	obj := &TestValidateRequiredStruct{
		Foo: "123",
	}

	err := val.ValidateStruct(obj)

	var data map[string]string
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestValidateRequiredStruct.Bar"] != "Bar is required" {
		t.Log(err.Error())
		t.Error("validate required failed")
	}
}

func TestValidateRequired3(t *testing.T) {
	val := validator.New()

	id, _ := uuid.Parse("1423ee2f-0f83-4693-a617-e06cb007b35e")
	obj := &TestValidateRequiredStruct{
		Foo: "123",
		Bar: id,
	}

	err := val.ValidateStruct(obj)
	if err != nil {
		t.Log(err.Error())
		t.Error("error must be nil")
	}
}

func TestValidateOnNil(t *testing.T) {
	val := validator.New()

	err := val.ValidateStruct(nil)
	if err != nil {
		t.Error(err.Error())
	}
}
