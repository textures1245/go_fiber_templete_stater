package apperror

import "errors"

var (
	ErrUserNotFound = errors.New("User not found while execute")
)
