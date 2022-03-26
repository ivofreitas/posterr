package follower

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strider-backend-test.com/api/routes/follower/model"
	"strider-backend-test.com/log"
)

type Controller interface {
	Follow(c echo.Context) error
	Unfollow(c echo.Context) error
}

type controller struct {
	service Service
	logger  *logrus.Entry
}

func NewController(service Service) Controller {
	return &controller{service, log.GetLogger()}
}

func (ctl *controller) Follow(c echo.Context) error {
	request := new(model.Request)

	if err := c.Bind(request); err != nil {
		ctl.logger.Warnf("error while trying bind: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(request); err != nil {
		ctl.logger.Warnf("error while trying validate: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	follower := new(model.Follower)
	follower.ID = request.Follower
	follower.Follow = request.Follow
	err := ctl.service.Follow(c.Request().Context(), follower)
	if err != nil {
		ctl.logger.Warnf("error while trying follow: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, follower)
}

func (ctl *controller) Unfollow(c echo.Context) error {
	request := new(model.Request)

	if err := c.Bind(request); err != nil {
		ctl.logger.Warnf("error while trying bind: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(request); err != nil {
		ctl.logger.Warnf("error while trying validate: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	follower := new(model.Follower)
	follower.ID = request.Follower
	follower.Follow = request.Follow
	err := ctl.service.Unfollow(c.Request().Context(), follower)
	if err != nil {
		ctl.logger.Warnf("error while trying unfollow: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, follower)
}
