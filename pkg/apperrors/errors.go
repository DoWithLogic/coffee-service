package apperrors

import "github.com/pkg/errors"

const (
	ERR_NOT_FOUND        = "Data Not Found"
	ERR_INTERNAL_SERVER  = "Internal Server Error"
	ERR_INVALID_PASSWORD = "Invalid User Password"
	ERR_BAD_REQUEST      = "Bad Request"
)

var (
	ErrNotFound        = errors.New(ERR_NOT_FOUND)
	ErrInternalServer  = errors.New(ERR_INTERNAL_SERVER)
	ErrInvalidPassword = errors.New(ERR_INVALID_PASSWORD)
	ErrBadRequest      = errors.New(ERR_BAD_REQUEST)
)
