package model

import (
	"time"

	"github.com/opencars/grpc/pkg/common"
	"github.com/opencars/schema/vehicle"
)

const DateLayout = "2006-01-02"

type Registration struct {
	Brand          *string  `json:"brand,omitempty"`
	Capacity       *int     `json:"capacity,omitempty"`
	Color          string   `json:"color"`
	FirstRegDate   *string  `json:"first_reg_date,omitempty"`
	Date           *string  `json:"date,omitempty"`
	Fuel           *string  `json:"fuel,omitempty"`
	Kind           *string  `json:"kind,omitempty"`
	Body           *string  `json:"body,omitempty"`
	Year           int      `json:"year"`
	Model          *string  `json:"model,omitempty"`
	DocumentNumber string   `json:"-"`
	DocumentSeries string   `json:"-"`
	Code           string   `json:"code"`
	Number         string   `json:"number"`
	NumSeating     *int     `json:"num_seating,omitempty"`
	NumStanding    *int     `json:"num_standing,omitempty"`
	OwnWeight      *float64 `json:"own_weight,omitempty"`
	RankCategory   *string  `json:"rank_category,omitempty"`
	TotalWeight    *float64 `json:"total_weight,omitempty"`
	VIN            *string  `json:"vin,omitempty"`
	IsActive       *bool    `json:"is_active,omitempty"`
}

func (r *Registration) Schema() *vehicle.Registration {
	item := vehicle.Registration{
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
		date, _ := time.Parse(DateLayout, *r.Date)
		item.Date = &common.Date{
			Year:  int32(date.Year()),
			Month: int32(date.Month()),
			Day:   int32(date.Day()),
		}
	}

	if r.FirstRegDate != nil {
		date, _ := time.Parse(DateLayout, *r.FirstRegDate)
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
