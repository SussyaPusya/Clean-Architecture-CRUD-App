package errogo

import "errors"

var (
	ErrUnauth = errors.New("Unauth")
	ErrAuth   = errors.New("you are Auth")
)
