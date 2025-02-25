package domain

import "errors"

var (
	ErrUserConflict       = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
