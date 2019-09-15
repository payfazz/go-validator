package validator_test

import (
	"encoding/json"
	"fmt"
	"github.com/payfazz/go-validator/validator"

	validator_v9 "gopkg.in/go-playground/validator.v9"

	"strings"
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

func TestRegisterValidation(t *testing.T) {
	type Test struct {
		Image string `validate:"type=jpg|type=png"`
	}

	val := validator.New()
	_ = val.Validate.RegisterValidation("type", func(f validator_v9.FieldLevel) bool {
		return strings.HasSuffix(f.Field().String(), f.Param())
	})

	err := val.ValidateStruct(Test{
		Image: "test.xyz",
	})

	fmt.Println(err)
}

func TestValidateOnNil(t *testing.T) {
	val := validator.New()

	err := val.ValidateStruct(nil)
	if err != nil {
		t.Error(err.Error())
	}
}
