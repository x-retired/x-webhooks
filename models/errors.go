package models

import "errors"

var (
	// UserNotExist: User does not Exit.
	UserNotExist = errors.New("User does not Exit.")
	// UserPasswordError: User password Exception.
	UserPasswordError = errors.New("User password Exception.")

	// InvalidParameter Invalid Parameter.
	InvalidParameter = errors.New("Invalid Parameter.")
)
