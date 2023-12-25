package apperrors

import "github.com/pkg/errors"

const (
	ERR_NOT_FOUND           = "Data Not Found"
	ERR_INTERNAL_SERVER     = "Internal Server Error"
	ERR_INVALID_PASSWORD    = "Invalid User Password"
	ERR_BAD_REQUEST         = "Bad Request"
	ERR_INVALID_CATEGORY_ID = "Invalid Menu Categories ID"
)

var (
	ErrNotFound                = errors.New(ERR_NOT_FOUND)
	ErrInternalServer          = errors.New(ERR_INTERNAL_SERVER)
	ErrInvalidPassword         = errors.New(ERR_INVALID_PASSWORD)
	ErrInvalidMenuCategoriesID = errors.New(ERR_INVALID_CATEGORY_ID)
	ErrBadRequest              = errors.New(ERR_BAD_REQUEST)
)
