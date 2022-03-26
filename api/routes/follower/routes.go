package follower

import (
	"github.com/labstack/echo/v4"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/follower/storage"
	"strider-backend-test.com/api/routes/user"
	userstorage "strider-backend-test.com/api/routes/user/storage"
)

func RouteMapping(group *echo.Group) {

	repository := mysql.New(mysql.GetConn())
	storage := storage.New(repository)

	userStorage := userstorage.New(repository)
	userService := user.NewService(userStorage)

	service := NewService(userService, storage)
	controller := NewController(service)

	group.POST("", controller.Follow)
	group.DELETE("", controller.Unfollow)
}
