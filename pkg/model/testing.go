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
		FirstRegDate: nil,
		Date:         nil,
		Fuel:         nil,
		Kind:         nil,
		Year:         0,
		Model:        nil,
		NDoc:         "000019",
		SDoc:         "CXE",
		Number:       "BC2425IC",
		NumSeating:   nil,
		NumStanding:  nil,
		OwnWeight:    nil,
		RankCategory: nil,
		TotalWeight:  nil,
		VIN:          nil,
	}
}
