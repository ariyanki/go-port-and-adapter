package error

import "errors"

var (
	InvalidUserNameOrPassword = errors.New("Invalid username or password")
	UserNotActive = errors.New("User is not active")
)