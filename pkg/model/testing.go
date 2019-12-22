package model

import (
	"testing"
)

// TODO: Implementation.
func TestRegistration(t *testing.T) *Registration {
	t.Helper()

	return &Registration{
		Brand:        nil,
		Capacity:     nil,
		Color:        "",
		DFirstReg:    nil,
		DReg:         nil,
		Fuel:         nil,
		Kind:         nil,
		MakeYear:     0,
		Model:        nil,
		NDoc:         "",
		SDoc:         "",
		NRegNew:      "",
		NSeating:     nil,
		NStanding:    nil,
		OwnWeight:    nil,
		RankCategory: nil,
		TotalWeight:  nil,
		VIN:          nil,
	}
}
