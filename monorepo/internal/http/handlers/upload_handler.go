package handlers

import (
	"net/http"
	"sylvie/internal/http/controllers"
	"sylvie/internal/http/dtos/response"
	"sylvie/internal/queue"

	"github.com/labstack/echo/v5"
)

func UploadHandler(
	uploadController controllers.UploadController,
	publisher queue.Publisher,
) echo.HandlerFunc {
	return func(c *echo.Context) error {
		title := c.FormValue("title")
		if title == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "title is required",
			})
		}

		file, err := c.FormFile("video")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "missing video",
			})
		}

		result, err := uploadController.Upload(file, title)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "upload failed",
			})
		}

		job := queue.Job{VideoID: result.VideoID, Path: result.Path}
		if err := publisher.Publish(job); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": "failed to publish upload event",
			})
		}

		return c.JSON(http.StatusCreated, response.UploadResponse{
			VideoID: result.VideoID,
			Status:  result.Status,
		})
	}
}
