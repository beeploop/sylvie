package handlers

import (
	"net/http"
	"slices"
	"strings"
	"sylvie/internal/config"
	"sylvie/internal/http/controllers"
	"sylvie/internal/http/views/pages"
	"sylvie/internal/utils"
	"sylvie/internal/video/entities"

	"github.com/labstack/echo/v5"
)

type HomepageQueryParams struct {
	Search string `query:"search"`
}

func Homepage(videosController controllers.VideosController) echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()

		var queries HomepageQueryParams
		if err := c.Bind(&queries); err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error="+err.Error())
		}

		videos, err := videosController.Search(queries.Search)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/?error="+err.Error())
		}

		viewmodel := pages.HomepageViewModel{
			Videos: slices.Collect(
				utils.Map(videos, func(video entities.Video) pages.Video {
					return pages.Video{
						ID:            video.ID,
						Title:         video.Title,
						ThumbnailPath: strings.TrimPrefix(video.ThumbnailPath, config.Load().Storage.BaseDir),
					}
				}),
			),
		}

		return pages.HomePage(viewmodel).Render(ctx, c.Response())
	}
}
