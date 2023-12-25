package http

import (
	jwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (h *handler) MapRoutes(echo *echo.Group) {
	echo.POST("/products/menu/categories", h.NewMenuCategoryHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.GET("/products/menu/categories/:id", h.DetailMenuCategoryHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.PUT("/products/menu/categories/:id", h.UpdateMenuCategoryHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.GET("/products/menu/categories/list", h.ListMenuCategoryHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))

	echo.POST("/products/menu/categories/:menu_categories_id", h.NewMenuHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.GET("/products/menu/:id", h.DetailMenuHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.PUT("/products/menu/:id", h.UpdateMenuHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
	echo.GET("/products/menu/list", h.ListMenuHandler, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
}
