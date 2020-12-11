package domain

type Error struct {
	msg string
}

func NewError(msg string) Error {
	return Error{
		msg: msg,
	}
}

func (e Error) Error() string {
	return e.msg
}

var (
	ErrNotFound  = NewError("record not found")
	ErrBadNumber = NewError("bad number")
	ErrBadVIN    = NewError("bad vin")
	ErrBadCode   = NewError("bad code")
)
