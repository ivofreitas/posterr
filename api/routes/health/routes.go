package health

import "github.com/labstack/echo/v4"

func RouteMapping(group *echo.Group) {
	group.GET("", health)
}
