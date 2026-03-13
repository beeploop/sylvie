package router

import (
	"sylvie/internal/application"
	"sylvie/internal/router/handlers"

	"github.com/labstack/echo/v5"
)

func RegisterRoutes(e *echo.Echo, app *application.Application) *echo.Echo {
	e.RouteNotFound("*", handlers.NotFoundHandler())

	e.POST("/upload", handlers.UploadHandler(app.UploadController))

	return e
}
