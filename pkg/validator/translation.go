package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

func RegisterTranslation(validate *validator.Validate, translator ut.Translator, tag string, translation string) {
	tags, translation := getAndReplaceTranslationKeywords(translation)

	err := validate.RegisterTranslation(tag,
		translator,
		func(ut ut.Translator) (err error) {
			if err = ut.Add(tag, translation, true); err != nil {
				return
			}
			return
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T(fe.Tag(), getParamByTags(tags, fe)...)
			if err != nil {
				return fe.(error).Error()
			}
			return t
		},
	)
	if err != nil {
		panic(err)
	}
}

func getParamByTags(tags []string, fe validator.FieldError) []string {
	var result []string

	for _, tag := range tags {
		switch tag {
		case "tag":
			result = append(result, fe.Tag())
		case "actualTag":
			result = append(result, fe.ActualTag())
		case "namespace":
			result = append(result, fe.Namespace())
		case "structNamespace":
			result = append(result, fe.StructNamespace())
		case "field":
			result = append(result, fe.Field())
		case "structField":
			result = append(result, fe.StructField())
		case "value":
			result = append(result, reflect.ValueOf(fe.Value()).String())
		case "param":
			result = append(result, fe.Param())
		}
	}

	return result
}

func getAndReplaceTranslationKeywords(s string) (tags []string, replaced string) {
	keywords := []string{
		"tag",
		"actualTag",
		"namespace",
		"structNamespace",
		"field",
		"structField",
		"value",
		"param",
	}

	regexString := ""
	for _, keyword := range keywords {
		if regexString != "" {
			regexString += "|"
		}
		regexString += "{" + keyword + "}"
	}

	re := regexp.MustCompile(regexString)
	tags = re.FindAllString(s, -1)
	for i := range tags {
		tags[i] = strings.ReplaceAll(tags[i], "{", "")
		tags[i] = strings.ReplaceAll(tags[i], "}", "")
	}

	i := -1
	replaced = re.ReplaceAllStringFunc(s,
		func(s string) string {
			i++
			return fmt.Sprintf("{%d}", i)
		},
	)

	return
}
