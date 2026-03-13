package handlers

import (
	"net/http"
	"sylvie/internal/router/controllers"

	"github.com/labstack/echo/v5"
)

func UploadHandler(uploadController controllers.UploadController) echo.HandlerFunc {
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

		return c.JSON(http.StatusCreated, result)
	}
}
