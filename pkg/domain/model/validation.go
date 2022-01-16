package model

import (
	"fmt"
	"strings"
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	Required     = "required"
	Invalid      = "invalid"
	IsNotInreger = "is_not_integer"
)

func Validate(v validation.Validatable, prefix string) error {
	err := v.Validate()
	if err == nil {
		return nil
	}

	messages := ErrorMessages(prefix, err)
	return NewValidationError(messages)
}

func ErrorMessages(prefix string, err error) map[string][]string {
	errs, ok := err.(validation.Errors)
	if !ok {
		return nil
	}

	messages := make(map[string][]string)
	for k, v := range errs {
		if _, ok := v.(validation.Errors); ok {
			key := fmt.Sprintf("%s.%s", prefix, ToSnakeCase(k))
			errMap := ErrorMessages(key, v)
			for k, v := range errMap {
				messages[ToSnakeCase(k)] = v
			}
		} else {
			key := fmt.Sprintf("%s.%s", prefix, ToSnakeCase(k))
			messages[key] = append(messages[key], v.Error())
		}
	}

	return messages
}

func ToSnakeCase(str string) string {
	var res = make([]rune, 0, len(str))
	for i, r := range str {
		if unicode.IsUpper(r) && i > 0 && (i+1 != len(str) && unicode.IsLower(rune(str[i+1]))) {
			res = append(res, '_', unicode.ToLower(r))
		} else {
			res = append(res, unicode.ToLower(r))
		}
	}

	return strings.ToLower(string(res))
}
