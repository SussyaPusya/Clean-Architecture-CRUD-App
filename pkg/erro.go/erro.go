package errogo

import "errors"

var (
	ErrUnauth      = errors.New("Unauth")
	ErrAuth        = errors.New("you are Auth")
	ErrUsrNotFOund = errors.New("User not Found")
	ErrUsrREg      = errors.New("username already in use")
)
