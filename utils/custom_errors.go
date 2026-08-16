package utils

import "errors"

var ErrDuplicateUsername = errors.New("username already exists")
var ErrInvalidCredentials = errors.New("invalid username or password")
var ErrInvalidToken = errors.New("invalid token")
var ErrInvalidSigningMethod = errors.New("Invalid signing method")
var ErrRefreshTokenExpired = errors.New("refresh token has expired")
var ErrUsernameMissing = errors.New("username missing from context")
var ErrCannotViewProfile = errors.New("Not authorized to view this profile")
