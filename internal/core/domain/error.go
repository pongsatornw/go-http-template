package domain

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// Assign the error variable from the jwt package
var ErrTokenExpired = jwt.ErrTokenExpired

// var ErrTokenExpired = errors.New("token is expired")
var ErrTokenInvalid = errors.New("token is invalid")
var ErrNoAuthProvided = errors.New("no auth provided")
var ErrInvalidAuthProvided = errors.New("invalid auth provided")
var ErrVerifyAuth = errors.New("failed to verify auth")
