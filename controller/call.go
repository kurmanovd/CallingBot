package controller

import (
	"context"
	"net/http"

	"server/logger"

	"server/service"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// CallController ...
type CallController struct {
	ctx      context.Context
	services *service.Manager
	logger   *logger.Logger
}

// NewCall creates a new call controller.
func NewCall(ctx context.Context, services *service.Manager, logger *logger.Logger) *CallController {
	return &CallController{
		ctx:      ctx,
		services: services,
		logger:   logger,
	}
}

// Get returns 200 and start call
func (ctr *CallController) Get(ctx echo.Context) error {
	err := ctr.services.Call.GetCall(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get call"))
	}

	return ctx.JSON(http.StatusOK, "")
}
