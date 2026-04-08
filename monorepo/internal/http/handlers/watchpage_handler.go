package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sylvie/internal/config"
	"sylvie/internal/http/controllers"
	"sylvie/internal/http/views/pages"

	"github.com/labstack/echo/v5"
)

type WatchpageQueryParams struct {
	VideoID string `query:"video" validate:"required"`
}

func WatchPage(videosController controllers.VideosController) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()

		params := url.Values{}

		var queries WatchpageQueryParams
		if err := c.Bind(&queries); err != nil {
			params.Add("error", err.Error())
			return c.Redirect(http.StatusSeeOther, "/?"+params.Encode())
		}

		video, err := videosController.Get(queries.VideoID)
		if err != nil {
			params.Add("error", err.Error())
			return c.Redirect(http.StatusSeeOther, "/watch?"+params.Encode())
		}

		viewmodel := pages.WatchpageViewModel{
			VideoURL: fmt.Sprintf("/media/%s", strings.TrimPrefix(video.MasterPlaylistPath, config.Load().Storage.BaseDir)),
		}

		return pages.WatchPage(viewmodel).Render(ctx, c.Response())
	}
}
