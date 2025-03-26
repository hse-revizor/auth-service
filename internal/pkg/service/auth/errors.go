package auth

import (
	"errors"
)

var (
	ErrUserExists   = errors.New("github user already exists")
	ErrUserNotFound = errors.New("github user not found")
)
