package errors

import "fmt"

//service grpc
var(
	ErrIncorrectPassword = fmt.Errorf("incorrect password")
	ErrUnknownLogin = fmt.Errorf("unknown login")
	ErrLoginTaken = fmt.Errorf("login already taken")
)

//transport grpc
var(
	ErrInvalidLogin = fmt.Errorf("invalid login")
	ErrInvalidPassword = fmt.Errorf("invalid password")
)