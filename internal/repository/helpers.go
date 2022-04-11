package repository

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrDuplicateEmail = errors.New("duplicate email")
)
