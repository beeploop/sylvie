package handlers

import (
	"net/http"
	"sylvie/internal/http/controllers"
	"sylvie/internal/http/dtos/response"

	"github.com/labstack/echo/v5"
)

func VideosHandler(videosController controllers.VideosController) echo.HandlerFunc {
	return func(c *echo.Context) error {
		videoID := c.Param("id")

		video, err := videosController.Get(videoID)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]any{
				"error": "video not found",
			})
		}

		return c.JSON(http.StatusOK, response.VideoResponse{
			VideoID:  video.ID,
			Title:    video.Title,
			Duration: video.DurationSeconds,
			Status:   video.Status,
		})
	}
}
