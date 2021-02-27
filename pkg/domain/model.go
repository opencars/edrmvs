package domain

type Registration struct {
	Brand          *string  `json:"brand,omitempty"`
	Capacity       *int     `json:"capacity,omitempty"`
	Color          string   `json:"color"`
	FirstRegDate   *string  `json:"first_reg_date,omitempty"`
	Date           *string  `json:"date,omitempty"`
	Fuel           *string  `json:"fuel,omitempty"`
	Kind           *string  `json:"kind,omitempty"`
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
