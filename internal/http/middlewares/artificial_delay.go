package middlewares

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v5"
)

func ArtificalDelay() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			if c.QueryParam("debug") == "slow" {
				fmt.Println("slow response turned on")
				time.Sleep(2000 * time.Millisecond)
			}
			return next(c)
		}
	}
}
