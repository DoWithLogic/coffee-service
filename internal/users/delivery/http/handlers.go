package http

import (
	"github.com/DoWithLogic/coffee-service/config"
	"github.com/DoWithLogic/coffee-service/internal/users"
	"github.com/DoWithLogic/coffee-service/internal/users/dtos"
	"github.com/DoWithLogic/coffee-service/pkg/apperrors"
	"github.com/DoWithLogic/coffee-service/pkg/observability"
	"github.com/DoWithLogic/coffee-service/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type handler struct {
	uc  users.UseCases
	log *zerolog.Logger
	cfg config.Config
}

func NewHandlers(uc users.UseCases, cfg config.Config) *handler {
	return &handler{uc: uc, log: observability.NewZeroLogHook().Z(), cfg: cfg}
}

func (h *handler) UserSignUp(c echo.Context) error {
	var request dtos.Users
	if err := c.Bind(&request); err != nil {
		h.log.Err(err).Msg("[users][UserSignUp]")

		return response.ErrorBuilders(apperrors.ErrBadRequest).Send(c)
	}

	err := h.uc.SignUp(c.Request().Context(), &request)
	if err != nil {
		return response.ErrorBuilders(err).Send(c)
	}

	return response.SuccessBuilder(map[string]interface{}{"id": request.ID}).Send(c)
}

func (h *handler) UserSignIn(c echo.Context) error {
	var request dtos.UserSignInRequest
	if err := c.Bind(&request); err != nil {
		h.log.Err(err).Msg("[users][UserSignIn]")

		return response.ErrorBuilders(apperrors.ErrBadRequest).Send(c)
	}

	signInData, err := h.uc.SignIn(c.Request().Context(), request)
	if err != nil {
		return response.ErrorBuilders(err).Send(c)
	}

	return response.SuccessBuilder(signInData).Send(c)
}

func (h *handler) UserDetailByID(c echo.Context) error {
	var request = new(struct {
		UserID int64 `param:"id"`
	})

	if err := c.Bind(request); err != nil {
		h.log.Err(err).Msg("[users][UserDetailByID]")

		return response.ErrorBuilders(apperrors.ErrBadRequest).Send(c)
	}

	users, err := h.uc.UserDetail(c.Request().Context(), request.UserID)
	if err != nil {
		return response.ErrorBuilders(err).Send(c)
	}

	return response.SuccessBuilder(users).Send(c)
}
