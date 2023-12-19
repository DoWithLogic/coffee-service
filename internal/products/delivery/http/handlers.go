package http

import (
	"github.com/DoWithLogic/coffee-service/config"
	"github.com/DoWithLogic/coffee-service/internal/products"
	"github.com/DoWithLogic/coffee-service/internal/products/dtos"
	"github.com/DoWithLogic/coffee-service/pkg/apperrors"
	"github.com/DoWithLogic/coffee-service/pkg/observability"
	"github.com/DoWithLogic/coffee-service/pkg/response"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type handler struct {
	uc  products.Usecase
	log *zerolog.Logger
	cfg config.Config
}

func NewHandlers(uc products.Usecase, cfg config.Config) *handler {
	return &handler{uc: uc, cfg: cfg, log: observability.NewZeroLogHook().Z()}
}

func (h *handler) NewMenuCategoryHandler(c echo.Context) error {
	var request dtos.MenuCategory
	if err := c.Bind(&request); err != nil {
		h.log.Err(err).Msg("[products][NewMenuCategoryHandler]Bind")

		return response.ErrorBuilders(apperrors.ErrBadRequest).Send(c)
	}

	if err := request.Validate(); err != nil {
		h.log.Err(err).Msg("[products][NewMenuCategoryHandler]Validate")

		return response.ErrorBuilders(apperrors.ErrBadRequest).Send(c)
	}

	err := h.uc.CreateMenuCategory(c.Request().Context(), &request)
	if err != nil {
		return response.ErrorBuilders(err).Send(c)
	}

	return response.SuccessBuilder(map[string]interface{}{"id": request.ID}).Send(c)
}

func (h *handler) DetailMenuCategoryHandler(c echo.Context) error {
	var request struct {
		MenuCategoryID int64 `param:"id"`
	}
	if err := c.Bind(&request); err != nil {
		h.log.Err(err).Msg("[products][DetailMenuCategoryHandler]Bind")

		return response.ErrorBuilders(apperrors.ErrBadRequest).Send(c)
	}

	menuCategoryData, err := h.uc.DetailMenuCategory(c.Request().Context(), request.MenuCategoryID)
	if err != nil {
		return response.ErrorBuilders(err).Send(c)
	}

	return response.SuccessBuilder(menuCategoryData).Send(c)
}

func (h *handler) UpdateMenuCategoryHandler(c echo.Context) error {
	var request dtos.UpdateMenuCategoryRequest
	if err := c.Bind(&request); err != nil {
		h.log.Err(err).Msg("[products][UpdateMenuCategoryHandler]Bind")

		return response.ErrorBuilders(apperrors.ErrBadRequest).Send(c)
	}

	if err := request.Validate(); err != nil {
		h.log.Err(err).Msg("[products][UpdateMenuCategoryHandler]Validate")

		return response.ErrorBuilders(apperrors.ErrBadRequest).Send(c)
	}

	err := h.uc.UpdateMenuCategory(c.Request().Context(), request)
	if err != nil {
		return response.ErrorBuilders(err).Send(c)
	}

	return response.SuccessBuilder(nil).Send(c)
}

func (h *handler) ListMenuCategoryHandler(c echo.Context) error {
	menuCategories, err := h.uc.ListMenuCategory(c.Request().Context())
	if err != nil {
		return response.ErrorBuilders(err).Send(c)
	}

	return response.SuccessBuilder(menuCategories).Send(c)
}
