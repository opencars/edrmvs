package query

import (
	"strconv"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/opencars/schema"
	"github.com/opencars/schema/vehicle"
	"github.com/opencars/seedwork"
	"github.com/opencars/translit"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

type ListByVIN struct {
	UserID  string
	TokenID string
	VIN     string
	Limit   string
	Offset  string
}

func (q *ListByVIN) Prepare() {
	q.VIN = translit.ToLatin(strings.ToUpper(q.VIN))
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
			validation.Required.Error(seedwork.Required),
		),
		validation.Field(
			&q.TokenID,
			validation.Required.Error(seedwork.Required),
		),
		validation.Field(
			&q.VIN,
			validation.Required.Error(seedwork.Required),
			validation.Length(6, 18).Error(seedwork.Invalid),
		),
		validation.Field(
			&q.Limit,
			is.Int.Error(seedwork.IsNotInreger),
		),
		validation.Field(
			&q.Offset,
			is.Int.Error(seedwork.IsNotInreger),
		),
	)
}

func (q *ListByVIN) Event(registrations ...model.Registration) schema.Producable {
	msg := vehicle.RegistrationSearched{
		UserId:       q.UserID,
		TokenId:      q.TokenID,
		Vin:          q.VIN,
		ResultAmount: uint32(len(registrations)),
		SearchedAt:   timestamppb.New(time.Now().UTC()),
	}

	return schema.New(&source, &msg).Message(
		schema.WithSubject(schema.CustomerRegistrationActions),
	)
}
