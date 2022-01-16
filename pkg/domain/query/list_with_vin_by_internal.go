package query

import (
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/opencars/translit"
)

type ListWithVINByInternal struct {
	VIN    string
	Limit  string
	Offset string
}

func (q *ListWithVINByInternal) Prepare() {
	q.VIN = translit.ToLatin(strings.ToUpper(q.VIN))
}

func (q *ListWithVINByInternal) GetOffset() uint64 {
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

func (q *ListWithVINByInternal) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.VIN,
			validation.Required.Error("required"),
			validation.Length(6, 18).Error("invalid"),
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
