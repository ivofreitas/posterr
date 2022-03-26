package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func health(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
