package post

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"strider-backend-test.com/api/routes/post/model"
	"strider-backend-test.com/log"
)

type Controller interface {
	Create(c echo.Context) error
	List(c echo.Context) error
}

type controller struct {
	service Service
	logger  *logrus.Entry
}

func NewController(service Service) Controller {
	return &controller{service, log.GetLogger()}
}

func (ctl *controller) Create(c echo.Context) error {

	request := new(model.Request)

	if err := c.Bind(request); err != nil {
		ctl.logger.Warnf("error while trying bind: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(request); err != nil {
		ctl.logger.Warnf("error while trying validate: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	post, err := ctl.service.Create(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, post)
}

func (ctl *controller) List(c echo.Context) error {

	listRequest := new(model.ListRequest)

	if err := c.Bind(listRequest); err != nil {
		ctl.logger.Warnf("error while trying bind: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := c.Validate(listRequest); err != nil {
		ctl.logger.Warnf("error while trying validate: %s", err.Error())
		return c.JSON(http.StatusBadRequest, err)
	}

	posts, err := ctl.service.List(c.Request().Context(), listRequest)
	if err != nil {
		ctl.logger.Warnf("error while trying list: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(posts) == 0 {
		return c.JSON(http.StatusNotFound, errors.New("no posts have been found"))
	}

	return c.JSON(http.StatusOK, posts)
}
