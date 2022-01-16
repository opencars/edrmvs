package query

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/opencars/translit"
)

type DetailsWithCodeByInternal struct {
	Code string
}

func (q *DetailsWithCodeByInternal) Prepare() {
	q.Code = translit.ToLatin(strings.ToUpper(q.Code))
}

func (q *DetailsWithCodeByInternal) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.Code,
			validation.Required.Error("required"),
			validation.Length(9, 9).Error("invalid"),
		),
	)
}
