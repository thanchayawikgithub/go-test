package services

import "errors"

var (
	ErrZeroAmount = errors.New("amount cannot be zero")
	ErrRepository = errors.New("repository error")
)
