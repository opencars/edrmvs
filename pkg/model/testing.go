package model

import (
	"testing"
)

// TestRegistration returns example registration.
func TestRegistration(t *testing.T) *Registration {
	t.Helper()

	return &Registration{
		Brand:        nil,
		Capacity:     nil,
		Color:        "СІРИЙ",
		DFirstReg:    nil,
		DReg:         nil,
		Fuel:         nil,
		Kind:         nil,
		MakeYear:     0,
		Model:        nil,
		NDoc:         "000019",
		SDoc:         "CXE",
		NRegNew:      "BC2425IC",
		NSeating:     nil,
		NStanding:    nil,
		OwnWeight:    nil,
		RankCategory: nil,
		TotalWeight:  nil,
		VIN:          nil,
	}
}
