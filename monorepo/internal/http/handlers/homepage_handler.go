package handlers

import (
	"sylvie/internal/http/views/pages"

	"github.com/labstack/echo/v5"
)

func Homepage() echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()

		return pages.HomePage().Render(ctx, c.Response())
	}
}
