package model

import "github.com/opencars/seedwork"

var (
	ErrNotFound  = seedwork.NewError("operation.not_found")
	ErrBadNumber = seedwork.NewError("query.bad_number")
	ErrBadVIN    = seedwork.NewError("query.bad_vin")
	ErrBadCode   = seedwork.NewError("query.bad_code")
)
