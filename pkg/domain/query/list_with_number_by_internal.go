package query

import (
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/opencars/translit"
)

type ListWithNumberByInternal struct {
	Number string
	Limit  string
	Offset string
}

func (q *ListWithNumberByInternal) Prepare() {
	q.Number = translit.ToLatin(strings.ToUpper(q.Number))
}

func (q *ListWithNumberByInternal) GetOffset() uint64 {
	if q.Offset == "" {
		return 0
	}

	num, err := strconv.ParseInt(q.Offset, 10, 64)
	if err != nil {
		panic(err)
	}

	if num < 0 {
		return 10
	}

	return uint64(num)
}

func (q *ListWithNumberByInternal) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.Number,
			validation.Required.Error("required"),
			validation.Length(2, 18).Error("invalid"),
		),
		validation.Field(
			&q.Limit,
			is.Int.Error("is_not_integer"),
		),
		validation.Field(
			&q.Offset,
			is.Int.Error("is_not_integer"),
		),
	)
}
