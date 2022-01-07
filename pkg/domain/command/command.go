package command

import (
	"fmt"
	"strings"
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/opencars/edrmvs/pkg/domain/model"
)

type Validatable interface {
	Validate() error
}

type Command interface {
	Validatable
}

func Process(cmd Command) error {
	return Validate(cmd, "request")
}

func Validate(v Validatable, prefix string) error {
	err := v.Validate()
	if err == nil {
		return nil
	}

	messages := ErrorMessages(prefix, err)
	return model.NewValidationError(messages)
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
		if unicode.IsLower(r) && i > 0 && (i+1 != len(str) && unicode.IsUpper(rune(str[i+1]))) {
			res = append(res, unicode.ToLower(r), '_')
		} else {
			res = append(res, unicode.ToLower(r))
		}
	}

	return strings.ToLower(string(res))
}
