package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sylvie/internal/config"
	"sylvie/internal/http/controllers"
	"sylvie/internal/http/views/pages/watchpage"
	"sylvie/internal/video/models"

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

		if video.Status != string(models.STATUS_READY) {
			return watchpage.NotReadyPage().Render(ctx, c.Response())
		}

		viewmodel := watchpage.WatchpageViewModel{
			VideoURL: fmt.Sprintf("/media/%s", strings.TrimPrefix(video.MasterPlaylistPath, config.Load().Storage.BaseDir)),
			Title:    video.Title,
		}

		return watchpage.WatchPage(viewmodel).Render(ctx, c.Response())
	}
}
