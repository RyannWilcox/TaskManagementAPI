package utils

import "errors"

var ErrDuplicateUsername = errors.New("username already exists")
var ErrInvalidCredentials = errors.New("invalid username or password")
var ErrInvalidToken = errors.New("invalid token")
var ErrInvalidSigningMethod = errors.New("Invalid signing method")
