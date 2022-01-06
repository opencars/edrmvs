package query

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type DetailsWithCodeByInternal struct {
	Code string
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
