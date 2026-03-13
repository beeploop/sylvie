package http

import (
	"sylvie/internal/application"
	"sylvie/internal/http/handlers"

	"github.com/labstack/echo/v5"
)

func registerRoutes(e *echo.Echo, app *application.Application) *echo.Echo {
	e.RouteNotFound("*", handlers.NotFoundHandler())

	e.POST("/uploads", handlers.UploadHandler(app.UploadController, app.Publisher))

	return e
}
