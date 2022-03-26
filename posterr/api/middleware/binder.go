package middleware

import (
	"github.com/labstack/echo/v4"
	"strings"
)

type Binder struct{}

func NewBinder() *Binder {
	return &Binder{}
}

func (cb *Binder) Bind(i interface{}, c echo.Context) (err error) {

	switch c.Request().Method {
	case "GET":
		qs := c.QueryString()

		var params []string
		params = append(params, qs)

		if !strings.Contains(qs, "offset") {
			params = append(params, "offset=0")
		}

		if !strings.Contains(qs, "limit") {
			params = append(params, "limit=10")
		}

		c.Request().URL.RawQuery = strings.Join(params, "&")
	}

	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != nil {
		return err
	}

	return
}
