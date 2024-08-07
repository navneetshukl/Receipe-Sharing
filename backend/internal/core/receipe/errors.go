package receipe

import "errors"

var (
	ErrAddingReceipe = errors.New("error in adding the receipe")
	ErrInvalidUserID = errors.New("invalid userid")
	ErrMissingField  = errors.New("some fields are missing")
)
