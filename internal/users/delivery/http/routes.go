package http

import (
	jwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (h *handler) MapRoutes(echo *echo.Group) {
	echo.POST("/users", h.UserSignUp)
	echo.POST("/users/login", h.UserSignIn)
	echo.GET("/users/:id", h.UserDetailByID, jwt.WithConfig(jwt.Config{SigningKey: []byte(h.cfg.Authorization.SecretKey)}))
}
