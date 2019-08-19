package validation_test

import (
	"testing"

	"github.com/payfazz/go-validator/pkg/validator"
)

var val *validator.Validator

func init() {
	val = validator.New()
}

type StructFoo struct {
	Date string `validate:"date_rfc3339"`
}

func TestDateRFC3339(t *testing.T) {
	foo := &StructFoo{
		Date: "2019-01-02T15:04:05+00:00",
	}
	err := val.ValidateStruct(foo)
	if nil != err {
		t.Fatal(err)
	}
}
