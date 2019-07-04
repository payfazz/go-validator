package validator_test

import (
	"testing"

	"github.com/payfazz/go-validator/pkg/validator"
)

type TestStruct struct {
	Foo string `validate:"required"`
}

func TestValidateRequired(t *testing.T) {
	val := validator.New()

	obj := &TestStruct{}

	err := val.ValidateStruct(obj)

	if err.Error() != "Foo is required" {
		t.Log(err.Error())
		t.Error("validate required failed")
	}
}
