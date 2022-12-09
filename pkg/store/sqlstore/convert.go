package sqlstore

import "github.com/opencars/edrmvs/pkg/domain/model"

func convertToDomain(from *Registration) *model.Registration {
	result := model.Registration{
		Brand:          from.Brand,
		Capacity:       from.Capacity,
		Color:          from.Color,
		FirstRegDate:   from.FirstRegDate,
		Date:           from.Date,
		Fuel:           from.Fuel,
		Kind:           from.Kind,
		Body:           from.Body,
		Year:           from.Year,
		Model:          from.Model,
		DocumentNumber: from.NDoc,
		DocumentSeries: from.SDoc,
		Code:           from.Code,
		Number:         from.Number,
		NumSeating:     from.NumSeating,
		NumStanding:    from.NumStanding,
		OwnWeight:      from.OwnWeight,
		RankCategory:   from.RankCategory,
		TotalWeight:    from.TotalWeight,
		VIN:            from.VIN,
	}

	if result.Date != nil {
		*result.Date = (*from.Date)[:10]
	}

	if result.FirstRegDate != nil {
		*result.FirstRegDate = (*from.FirstRegDate)[:10]
	}

	return &result
}

func convertFromDomain(from *model.Registration) *Registration {
	result := Registration{
		Brand:        from.Brand,
		Capacity:     from.Capacity,
		Color:        from.Color,
		FirstRegDate: from.FirstRegDate,
		Date:         from.Date,
		Fuel:         from.Fuel,
		Kind:         from.Kind,
		Year:         from.Year,
		Model:        from.Model,
		NDoc:         from.DocumentNumber,
		SDoc:         from.DocumentSeries,
		Code:         from.Code,
		Number:       from.Number,
		NumSeating:   from.NumSeating,
		NumStanding:  from.NumStanding,
		OwnWeight:    from.OwnWeight,
		RankCategory: from.RankCategory,
		TotalWeight:  from.TotalWeight,
		VIN:          from.VIN,
	}

	return &result
}
