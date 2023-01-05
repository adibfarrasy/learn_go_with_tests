package banking

import "github.com/pkg/errors"

var (
	Unknown             = errors.New("unknown error")
	InsufficientBalance = errors.New("insufficient balance")
)
