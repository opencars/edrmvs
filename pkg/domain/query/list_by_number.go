package query

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/opencars/schema"
	"github.com/opencars/schema/vehicle"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

type ListByNumber struct {
	UserID string
	Number string
	Limit  string
	Offset string
}

func (q *ListByNumber) GetLimit() uint64 {
	if q.Limit == "" {
		return 10
	}

	num, err := strconv.ParseInt(q.Limit, 10, 64)
	if err != nil {
		panic(err)
	}

	if num < 0 {
		return 10
	}

	return uint64(num)
}

func (q *ListByNumber) GetOffset() uint64 {
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

func (q *ListByNumber) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.UserID,
			validation.Required.Error("required"),
		),
		validation.Field(
			&q.Number,
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

func (q *ListByNumber) Event(registrations ...model.Registration) schema.Producable {
	msg := vehicle.RegistrationSearched{
		UserId:       q.UserID,
		Number:       q.Number,
		ResultAmount: uint32(len(registrations)),
	}

	return schema.New(&source, &msg).Message(
		schema.WithSubject(schema.CustomerRegistrationActions),
	)
}
