package post

import (
	"github.com/labstack/echo/v4"
	"strider-backend-test.com/adapter/mysql"
	"strider-backend-test.com/api/routes/post/storage"
	"strider-backend-test.com/api/routes/user"
	userstorage "strider-backend-test.com/api/routes/user/storage"
)

func RouteMapping(group *echo.Group) {

	conn := mysql.GetConn()
	repository := mysql.New(conn)
	storage := storage.New(repository)

	userStorage := userstorage.New(repository)
	userService := user.NewService(userStorage)

	service := NewService(userService, storage)
	controller := NewController(service)

	group.POST("", controller.Create)
	group.GET("", controller.List)
	group.POST("/quote", controller.Create)
	group.POST("/repost", controller.Create)
}
