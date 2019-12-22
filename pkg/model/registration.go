package model

import (
	"fmt"
	"strconv"

	"github.com/opencars/edrmvs/pkg/hsc"
	"github.com/opencars/edrmvs/pkg/utils"
)

// Registration represents database entity for vehicle registration.
type Registration struct {
	Brand        *string `db:"brand" json:"brand"`
	Capacity     *int    `db:"capacity" json:"capacity"`
	Color        string  `db:"color" json:"color"`
	DFirstReg    *string `db:"d_first_reg" json:"d_first_reg"`
	DReg         *string `db:"d_reg" json:"d_reg"`
	Fuel         *string `db:"fuel" json:"fuel"`
	Kind         *string `db:"kind" json:"kind"`
	MakeYear     int     `db:"make_year" json:"make_year"`
	Model        *string `db:"model" json:"model"`
	NDoc         string  `db:"n_doc" json:"n_doc"`
	SDoc         string  `db:"s_doc" json:"s_doc"`
	NRegNew      string  `db:"n_reg_new" json:"n_reg_new"`
	NSeating     *int    `db:"n_seating" json:"n_seating"`
	NStanding    *int    `db:"n_standing" json:"n_standing"`
	OwnWeight    *int    `db:"own_weight" json:"own_weight"`
	RankCategory *string `db:"rank_category" json:"rank_category"`
	TotalWeight  *int    `db:"total_weight" json:"total_weight"`
	VIN          *string `db:"vin" json:"vin"`
}

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

	ownWeight, err := utils.Atoi(registration.OwnWeight)
	if err != nil {
		return nil, fmt.Errorf("ownWeight: %w", err)
	}

	totalWeight, err := utils.Atoi(registration.TotalWeight)
	if err != nil {
		return nil, fmt.Errorf("totalWeight: %w", err)
	}

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
		SDoc:         registration.SDoc,
		NRegNew:      registration.NRegNew,
		NSeating:     nSeating,
		NStanding:    nStanding,
		OwnWeight:    ownWeight,
		RankCategory: registration.RankCategory,
		TotalWeight:  totalWeight,
		VIN:          registration.VIN,
	}, nil
}
