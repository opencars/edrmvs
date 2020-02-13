package model

import (
	"fmt"
	"strconv"

	"github.com/opencars/edrmvs/pkg/hsc"
	"github.com/opencars/edrmvs/pkg/utils"
	"github.com/opencars/translit"
)

// Registration represents database entity for vehicle registration.
type Registration struct {
	Brand        *string  `db:"brand" json:"brand"`
	Capacity     *int     `db:"capacity" json:"capacity"`
	Color        string   `db:"color" json:"color"`
	DFirstReg    *string  `db:"d_first_reg" json:"first_reg_date"`
	DReg         *string  `db:"d_reg" json:"date"`
	Fuel         *string  `db:"fuel" json:"fuel"`
	Kind         *string  `db:"kind" json:"kind"`
	MakeYear     int      `db:"make_year" json:"year"`
	Model        *string  `db:"model" json:"model"`
	NDoc         string   `db:"n_doc" json:"-"`
	SDoc         string   `db:"s_doc" json:"-"`
	Code         string   `db:"-" json:"code"`
	NRegNew      string   `db:"n_reg_new" json:"number"`
	NSeating     *int     `db:"n_seating" json:"num_seating"`
	NStanding    *int     `db:"n_standing" json:"num_standing"`
	OwnWeight    *float64 `db:"own_weight" json:"own_weight"`
	RankCategory *string  `db:"rank_category" json:"rank_category"`
	TotalWeight  *float64 `db:"total_weight" json:"total_weight"`
	VIN          *string  `db:"vin" json:"vin"`
}

// FromHSC returns parsed and prettified registration.
func FromHSC(registration hsc.Registration) (*Registration, error) {
	capacity, err := utils.Atoi(registration.Capacity)
	if err != nil {
		return nil, fmt.Errorf("capacity: %w", err)
	}

	makeYear, err := strconv.Atoi(registration.MakeYear)
	if err != nil {
		return nil, fmt.Errorf("makeYear: %w", err)
	}

	nSeating, err := utils.Atoi(registration.NSeating)
	if err != nil {
		return nil, fmt.Errorf("nSeating: %w", err)
	}

	nStanding, err := utils.Atoi(registration.NStanding)
	if err != nil {
		return nil, fmt.Errorf("nStanding: %w", err)
	}

	ownWeight, err := utils.Atof(registration.OwnWeight)
	if err != nil {
		return nil, fmt.Errorf("ownWeight: %w", err)
	}

	totalWeight, err := utils.Atof(registration.TotalWeight)
	if err != nil {
		return nil, fmt.Errorf("totalWeight: %w", err)
	}

	translit.ToLatin(registration.SDoc)

	return &Registration{
		Brand:        registration.Brand,
		Capacity:     capacity,
		Color:        registration.Color,
		DFirstReg:    registration.DFirstReg,
		DReg:         registration.DReg,
		Fuel:         registration.Fuel,
		Kind:         registration.Kind,
		MakeYear:     makeYear,
		Model:        registration.Model,
		NDoc:         registration.NDoc,
		SDoc:         translit.ToLatin(registration.SDoc),
		NRegNew:      translit.ToLatin(registration.NRegNew),
		NSeating:     nSeating,
		NStanding:    nStanding,
		OwnWeight:    ownWeight,
		RankCategory: registration.RankCategory,
		TotalWeight:  totalWeight,
		VIN:          registration.VIN,
	}, nil
}
