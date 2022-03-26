package routes

import (
	"github.com/labstack/echo/v4"
	"strider-backend-test.com/api/routes/follower"
	"strider-backend-test.com/api/routes/health"
	"strider-backend-test.com/api/routes/post"
)

type FnRouteMapping func(r *echo.Group)

func RouteMapping(group *echo.Group) {
	route(group, "/posts", post.RouteMapping)
	route(group, "/health", health.RouteMapping)
	route(group, "/followers", follower.RouteMapping)
}

func route(r *echo.Group, prefix string, fn FnRouteMapping) {
	fn(r.Group(prefix))
}
