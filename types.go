package uhoh

import (
	"errors"
)

var (
	ErrGeneral    = errors.New("general error")
	ErrBadRequest = errors.New("bad request error")
	ErrNotFound   = errors.New("not found error")

	// Data checking
	ErrValidation = errors.New("validation error")

	// Permissions
	ErrForbidden    = errors.New("forbidden error")
	ErrPermission   = errors.New("permission error")
	ErrUnauthorized = errors.New("unauthorized error")

	// Database
	ErrDBConnection = errors.New("database connection error")
	ErrDBQuery      = errors.New("database query error")
	ErrDBNoRows     = errors.New("database no rows error")

	// Internal errors
	ErrInternal = errors.New("internal error")
)
