package grpc

import (
	"time"

	"github.com/opencars/grpc/pkg/common"
	"github.com/opencars/grpc/pkg/registration"

	"github.com/opencars/edrmvs/pkg/domain/model"
)

func FromDomain(r *model.Registration) *registration.Record {
	item := registration.Record{
		Code:   r.Code,
		Number: r.Number,
		Year:   int32(r.Year),
		Color:  r.Color,
	}

	if r.VIN != nil {
		item.Vin = *r.VIN
	}

	if r.Brand != nil {
		item.Brand = *r.Brand
	}

	if r.Model != nil {
		item.Model = *r.Model
	}

	if r.Capacity != nil {
		item.Capacity = int32(*r.Capacity)
	}

	if r.Fuel != nil {
		item.Fuel = *r.Fuel
	}

	if r.Kind != nil {
		item.Kind = *r.Kind
	}

	if r.NumSeating != nil {
		item.NumSeating = int32(*r.NumSeating)
	}

	if r.OwnWeight != nil {
		item.OwnWeight = int32(*r.OwnWeight)
	}

	if r.TotalWeight != nil {
		item.TotalWeight = int32(*r.TotalWeight)
	}

	if r.TotalWeight != nil {
		item.TotalWeight = int32(*r.TotalWeight)
	}

	if r.Date != nil {
		date, _ := time.Parse(model.DateLayout, *r.Date)
		item.Date = &common.Date{
			Year:  int32(date.Year()),
			Month: int32(date.Month()),
			Day:   int32(date.Day()),
		}
	}

	if r.FirstRegDate != nil {
		date, _ := time.Parse(model.DateLayout, *r.FirstRegDate)
		item.FirstRegDate = &common.Date{
			Year:  int32(date.Year()),
			Month: int32(date.Month()),
			Day:   int32(date.Day()),
		}
	}

	// TODO: Convert to enum. (changes required in domain)
	if r.RankCategory != nil {
		item.Category = *r.RankCategory
	}

	return &item
}
