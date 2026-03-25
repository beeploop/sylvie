package handlers

import (
	"fmt"
	"net/http"
	"sylvie/internal/http/views/pages"

	"github.com/labstack/echo/v5"
)

type HomepageQueryParams struct {
	Search string `query:"search"`
}

func Homepage() echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()

		var queries HomepageQueryParams
		if err := c.Bind(&queries); err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error="+err.Error())
		}

		fmt.Println(queries)

		return pages.HomePage().Render(ctx, c.Response())
	}
}
