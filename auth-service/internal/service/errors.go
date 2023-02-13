package service

import "errors"

var (
	ErrUserExists   = errors.New("email is busy")
	ErrUserNotExist = errors.New("user with this email not found")
)
