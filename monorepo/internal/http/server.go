package http

import (
	"net/http"
	"sylvie/internal/application"
	"sylvie/internal/config"

	"github.com/labstack/echo/v5"
)

func NewServer(app *application.Application) *http.Server {
	r := echo.New()

	return &http.Server{
		Addr:    config.Load().Server.PORT,
		Handler: registerRoutes(r, app),
	}
}
