package handlers

import (
	"net/http"
	"net/url"
	"sylvie/internal/http/controllers"
	"sylvie/internal/queue"

	"github.com/labstack/echo/v5"
)

func UploadHandler(
	uploadController controllers.UploadController,
	publisher queue.Publisher,
) echo.HandlerFunc {
	return func(c *echo.Context) error {
		params := url.Values{}
		u, _ := url.Parse("/")

		sourceForm := c.FormValue("source_dialog")
		params.Add("dialog", sourceForm)

		title := c.FormValue("title")
		if title == "" {
			params.Add("error", "title required")
			u.RawQuery = params.Encode()
			return c.Redirect(http.StatusSeeOther, u.String())
		}

		file, err := c.FormFile("video")
		if err != nil {
			params.Add("error", "title required")
			u.RawQuery = params.Encode()
			return c.Redirect(http.StatusSeeOther, u.String())
		}

		video, err := uploadController.Upload(file, title)
		if err != nil {
			params.Add("error", err.Error())
			u.RawQuery = params.Encode()
			return c.Redirect(http.StatusSeeOther, u.String())
		}

		job := queue.Job{VideoID: video.ID, Path: video.OriginalPath}
		if err := publisher.Publish(job); err != nil {
			params.Add("error", err.Error())
			u.RawQuery = params.Encode()
			return c.Redirect(http.StatusSeeOther, u.String())
		}

		return c.Redirect(http.StatusSeeOther, u.String())
	}
}
