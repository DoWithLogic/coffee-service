package response

import (
	"net/http"

	"github.com/DoWithLogic/coffee-service/pkg/apperrors"
	"github.com/labstack/echo/v4"
)

type CustomResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

var MapErrors = map[error]CustomResponse{
	apperrors.ErrNotFound:        {Code: http.StatusNotFound, Message: http.StatusText(http.StatusNotFound), Error: apperrors.ERR_NOT_FOUND},
	apperrors.ErrInvalidPassword: {Code: http.StatusBadRequest, Message: http.StatusText(http.StatusBadRequest), Error: apperrors.ERR_INVALID_PASSWORD},

	apperrors.ErrBadRequest:     {Code: http.StatusBadRequest, Message: apperrors.ERR_BAD_REQUEST, Error: apperrors.ERR_BAD_REQUEST},
	apperrors.ErrInternalServer: {Code: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError), Error: apperrors.ERR_INTERNAL_SERVER},
}

func (c CustomResponse) isNull() bool {
	return c == CustomResponse{}
}

func ErrorBuilders(err error) CustomResponse {
	customErr := MapErrors[err]
	if customErr.isNull() {
		return CustomResponse{
			Code:    http.StatusInternalServerError,
			Message: apperrors.ERR_INTERNAL_SERVER,
			Error:   err.Error(),
		}
	}

	return customErr
}

func SuccessBuilder(response interface{}) CustomResponse {
	return CustomResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    response,
	}
}

func (x CustomResponse) Send(c echo.Context) error {
	return c.JSON(x.Code, x)
}
