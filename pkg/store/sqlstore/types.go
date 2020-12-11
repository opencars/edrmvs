package sqlstore

type Registration struct {
	Brand        *string  `db:"brand"`
	Capacity     *int     `db:"capacity"`
	Color        string   `db:"color"`
	FirstRegDate *string  `db:"d_first_reg"`
	Date         *string  `db:"d_reg"`
	Fuel         *string  `db:"fuel"`
	Kind         *string  `db:"kind"`
	Year         int      `db:"make_year"`
	Model        *string  `db:"model"`
	NDoc         string   `db:"n_doc"`
	SDoc         string   `db:"s_doc"`
	Code         string   `db:"code"`
	Number       string   `db:"n_reg_new"`
	NumSeating   *int     `db:"n_seating"`
	NumStanding  *int     `db:"n_standing"`
	OwnWeight    *float64 `db:"own_weight"`
	RankCategory *string  `db:"rank_category"`
	TotalWeight  *float64 `db:"total_weight"`
	VIN          *string  `db:"vin"`
}
