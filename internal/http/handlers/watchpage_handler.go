package handlers

import (
	"net/http"
	"net/url"
	"path/filepath"
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

		relative, err := filepath.Rel(config.Load().Storage.BaseDir, video.MasterPlaylistPath)
		if err != nil {
			params.Add("error", err.Error())
			return c.Redirect(http.StatusSeeOther, "/watch?"+params.Encode())
		}

		viewmodel := watchpage.WatchpageViewModel{
			VideoURL: "/media/" + filepath.ToSlash(relative),
			Title:    video.Title,
		}

		return watchpage.WatchPage(viewmodel).Render(ctx, c.Response())
	}
}
