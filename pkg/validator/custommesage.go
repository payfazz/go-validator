package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
)

//Message is key-value pair of tag and translation
type Messages map[string]string

//RegisterMessages register translation with key-value pair of tag and translation string
//this method is not thread-safe it is intended that these all be registered prior to any validation
func (val *Validator) RegisterMessages(trans Messages) {
	for k, v := range trans {
		val.registerMessage(k, v)
	}
}

func (val *Validator) registerMessage(tag string, translation string) {
	validate := val.Validate
	translator := val.Trans

	tags, translation := getAndReplaceMessageKeywords(translation)

	err := validate.RegisterTranslation(tag,
		translator,
		func(ut ut.Translator) (err error) {
			if err = ut.Add(tag, translation, true); err != nil {
				return
			}
			return
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), getParamByTags(tags, fe)...)
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
			val := reflect.ValueOf(fe.Value())
			typ := val.Type().Name()
			var s string
			switch typ {
			case "float32":
				fallthrough
			case "float64":
				s = fmt.Sprintf("%.2f", val.Float())
			case "int":
				fallthrough
			case "int32":
				fallthrough
			case "int64":
				s = fmt.Sprintf("%d", val.Int())
			default:
				s = val.String()
			}
			result = append(result, s)
		case "param":
			result = append(result, fe.Param())
		}
	}

	return result
}

func getAndReplaceMessageKeywords(s string) (tags []string, replaced string) {
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

func getAndReplaceMessageKeywordsSprintf(s string) (tags []string, replaced string) {
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

	replaced = re.ReplaceAllStringFunc(s,
		func(s string) string {
			return fmt.Sprintf("%%s")
		},
	)

	return
}
