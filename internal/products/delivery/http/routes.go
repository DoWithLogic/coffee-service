package http

import (
	jwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (h *handler) MapRoutes(echo *echo.Group) {
	echo.POST("/products", h.NewMenuCategoryHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.GET("/products/:id", h.DetailMenuCategoryHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.PUT("/products/:id", h.UpdateMenuCategoryHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.GET("/products/list", h.ListMenuCategoryHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
}
