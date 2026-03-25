package http

import (
	"embed"
	"sylvie/internal/application"
	"sylvie/internal/http/handlers"

	"github.com/labstack/echo/v5"
)

//go:embed static/*
var assets embed.FS

func registerRoutes(e *echo.Echo, app *application.Application) *echo.Echo {
	e.RouteNotFound("*", handlers.NotFoundHandler())

	e.StaticFS("/styles", echo.MustSubFS(assets, "static/css"))
	e.StaticFS("/scripts", echo.MustSubFS(assets, "static/js"))
	e.StaticFS("/assets", echo.MustSubFS(assets, "static/assets"))

	e.GET("/", handlers.Homepage(app.VideosController))

	e.POST("/uploads", handlers.UploadHandler(app.UploadController, app.Publisher))
	e.GET("/videos/:id", handlers.VideosHandler(app.VideosController))

	return e
}
