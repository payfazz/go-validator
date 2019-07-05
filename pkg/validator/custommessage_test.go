package validator_test

import (
	"encoding/json"
	"testing"

	"github.com/payfazz/go-validator/pkg/validator"
)

func TestTranslateField(t *testing.T) {
	obj := &TestTagStruct{}

	custom := map[string]string{
		"A.min": "{field} minimal {param}!",
		"C.min": "{field} length please at least {param}!",
	}

	val := validator.New()

	err := val.WithCustomFieldMessages(custom).ValidateStruct(obj)

	var data map[string]string
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestTagStruct.A"] != "A minimal 5!" {
		t.Log(err)
		t.Error("wrong custom field translation")
	}

	err = val.ValidateStruct(obj)
	json.Unmarshal([]byte(err.Error()), &data)

	if data["TestTagStruct.A"] != "A must be minimal 5" {
		t.Log(err)
		t.Error("a")
	}

}
