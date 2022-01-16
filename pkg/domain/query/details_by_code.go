package query

import (
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/opencars/schema"
	"github.com/opencars/schema/vehicle"
	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

type DetailsByCode struct {
	UserID  string
	TokenID string
	Code    string
}

func (q *DetailsByCode) Prepare() {
	q.Code = translit.ToLatin(strings.ToUpper(q.Code))
}

func (q *DetailsByCode) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.UserID,
			validation.Required.Error(model.Required),
		),
		validation.Field(
			&q.TokenID,
			validation.Required.Error(model.Required),
		),
		validation.Field(
			&q.Code,
			validation.Required.Error(model.Required),
			validation.Length(9, 9).Error(model.Invalid),
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
		SearchedAt:   timestamppb.New(time.Now().UTC()),
	}

	return schema.New(&source, &msg).Message(
		schema.WithSubject(schema.CustomerRegistrationActions),
	)
}
