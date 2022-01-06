package query

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/opencars/schema"
	"github.com/opencars/schema/vehicle"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

type ListByVIN struct {
	UserID string
	VIN    string
	Limit  string
	Offset string
}

func (q *ListByVIN) GetLimit() uint64 {
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

func (q *ListByVIN) GetOffset() uint64 {
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

func (q *ListByVIN) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.UserID,
			validation.Required.Error("required"),
		),
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

func (q *ListByVIN) Event(registrations ...model.Registration) schema.Producable {
	msg := vehicle.RegistrationSearched{
		UserId:       q.UserID,
		Vin:          q.VIN,
		ResultAmount: uint32(len(registrations)),
	}

	return schema.NewMessage(&msg).WithOptions(
		schema.WithSubject(schema.RegistrationCustomerActions),
	)
}
