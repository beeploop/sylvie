package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func NotFoundHandler() echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.JSON(http.StatusNotFound, map[string]any{
			"error": "not found route",
		})
	}
}
