package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DoWithLogic/coffee-service/config"
	"github.com/DoWithLogic/coffee-service/pkg/databases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	httpProducts "github.com/DoWithLogic/coffee-service/internal/products/delivery/http"
	repoProducts "github.com/DoWithLogic/coffee-service/internal/products/repository"
	ucProducts "github.com/DoWithLogic/coffee-service/internal/products/usecase"

	httpUsers "github.com/DoWithLogic/coffee-service/internal/users/delivery/http"
	repoUsers "github.com/DoWithLogic/coffee-service/internal/users/repository"
	ucUsers "github.com/DoWithLogic/coffee-service/internal/users/usecase"
)

func main() {
	cfg, err := config.LoadConfig("config")
	if err != nil {
		panic(err)
	}

	// Set the time zone to Asia/Jakarta
	_, err = time.LoadLocation(cfg.Server.TimeZone)
	if err != nil {
		log.Error("Error on Set the time zone to Asia/Jakarta: ", err)

		return
	}

	db, err := databases.NewMySQLDB(context.Background(), cfg.Database)
	if err != nil {
		panic(err)
	}

	server := echo.New()
	server.Use(middleware.CORS())

	version := server.Group("/v1/coffee")

	version.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "welcome to coffe-service")
	})

	usersRepo := repoUsers.NewRepository(db)
	productsRepo := repoProducts.NewRepository(db)

	usersUC := ucUsers.NewUsecase(usersRepo, cfg)
	productsUC := ucProducts.NewUseCase(productsRepo)

	usersHandlers := httpUsers.NewHandlers(usersUC, cfg)
	productsHandlers := httpProducts.NewHandlers(productsUC, cfg)

	usersHandlers.MapRoutes(version)
	productsHandlers.MapRoutes(version)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	go func() {
		<-quit
		log.Info("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db.Close()
		server.Shutdown(ctx)
	}()

	log.Info(server.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
