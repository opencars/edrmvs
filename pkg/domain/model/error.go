package model

var (
	ErrNotFound  = NewError("operation.not_found")
	ErrBadNumber = NewError("query.bad_number")
	ErrBadVIN    = NewError("query.bad_vin")
	ErrBadCode   = NewError("query.bad_code")
)

type Error struct {
	text string
}

func NewError(text string) Error {
	return Error{
		text: text,
	}
}

func (e Error) Error() string {
	return e.text
}
