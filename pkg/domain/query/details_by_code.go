package query

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/opencars/schema"
	"github.com/opencars/schema/vehicle"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

type DetailsByCode struct {
	UserID string
	Code   string
}

func (q *DetailsByCode) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.UserID,
			validation.Required.Error("required"),
		),
		validation.Field(
			&q.Code,
			validation.Required.Error("required"),
			validation.Length(9, 9).Error("invalid"),
		),
	)
}

func (q *DetailsByCode) Event(registration *model.Registration) schema.Producable {
	var resultAmount uint32
	if registration != nil {
		resultAmount = 1
	}
	msg := vehicle.RegistrationSearched{
		UserId:       q.UserID,
		Code:         q.Code,
		ResultAmount: resultAmount,
	}

	return schema.New(&source, &msg).Message(
		schema.WithSubject(schema.RegistrationCustomerActions),
	)
}
