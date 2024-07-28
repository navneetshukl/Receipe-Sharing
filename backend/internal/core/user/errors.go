package user

import "errors"

var (
	ErrMissingField       = errors.New("some of the fields are not present")
	ErrInvalidPhoneNumber = errors.New("no phone number is present")
	ErrHashingPassword    = errors.New("error in hashing password")
	ErrAddingUser         = errors.New("error in inserting the user")
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrUserAlreadyExist   = errors.New("user already exist")
)
