package main

import (
	"context"
	"net/http"
	"time"

	"server/service"

	"github.com/pkg/errors"

	echoLog "github.com/labstack/gommon/log"

	"server/config"
	"server/controller"

	libError "server/lib/error"
	"server/lib/validator"
	"server/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	// config
	cfg := config.Get()

	// logger
	l := logger.Get()

	// Init service manager
	serviceManager, err := service.NewManager(ctx)
	if err != nil {
		return errors.Wrap(err, "manager.New failed")
	}

	// Init controllers
	callController := controller.NewCall(ctx, serviceManager, l)

	// Initialize Echo instance
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"content-type", http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		ExposeHeaders: []string{"Content-Range"},
	}))
	e.Validator = validator.NewValidator()
	e.HTTPErrorHandler = libError.Error
	// Disable Echo JSON logger in debug mode
	if cfg.LogLevel == "debug" {
		if l, ok := e.Logger.(*echoLog.Logger); ok {
			l.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
		}
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// API V1
	v1 := e.Group("/v1")

	// Call routes
	callRoutes := v1.Group("/call")
	callRoutes.GET("", callController.Get)

	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))

	return nil
}
