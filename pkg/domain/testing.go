package domain

import "testing"

// TestRegistration returns example registration.
func TestRegistration(t *testing.T) *Registration {
	t.Helper()

	brand := "TESLA"
	model := "MODEL X"
	date := "2019-06-05"
	fuel := "ЕЛЕКТРО"
	kind := "ЛЕГКОВИЙ УНІВЕРСАЛ-B"
	firstRegDate := "2016-10-13"
	vin := "5YJXCCE40GF010543"
	numSeating := 7
	ownWeight := 2485.0
	rankCategory := "B"
	totalWeight := 3021.0

	return &Registration{
		Brand:          &brand,
		Color:          "ЧОРНИЙ",
		FirstRegDate:   &firstRegDate,
		Date:           &date,
		Fuel:           &fuel,
		Kind:           &kind,
		Year:           2016,
		Model:          &model,
		DocumentNumber: "484154",
		DocumentSeries: "CXH",
		Code:           "CXH484154",
		Number:         "AA9359PC",
		NumSeating:     &numSeating,
		OwnWeight:      &ownWeight,
		RankCategory:   &rankCategory,
		TotalWeight:    &totalWeight,
		VIN:            &vin,
	}
}
