package api

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strider-backend-test.com/api/middleware"
	"strider-backend-test.com/api/routes"
	"strider-backend-test.com/config"
	"strider-backend-test.com/log"
	"syscall"
	"time"
)

type Server struct {
	http   *echo.Echo
	logger *logrus.Entry
	signal chan struct{}
}

func NewServer() *Server {
	return &Server{
		logger: log.GetLogger(),
		signal: make(chan struct{}),
	}
}

func (s *Server) Run() {
	s.start()
	s.logger.Println("Server started and waiting for the graceful signal...")
	<-s.signal
}

func (s *Server) start() {
	go s.watchStop()

	config := config.Get()
	serverConfig := config.Server

	s.http = echo.New()
	s.logger.Infof("Server is starting in port %s.", serverConfig.Port)

	s.http.Validator = middleware.NewValidator()
	s.http.Binder = middleware.NewBinder()
	s.http.Use(echomiddleware.Logger())
	s.http.Use(echomiddleware.Recover())
	s.http.Pre(echomiddleware.RemoveTrailingSlash())

	r := s.http.Group(serverConfig.BasePath)
	routes.RouteMapping(r)

	addr := fmt.Sprintf(":%s", serverConfig.Port)
	go func() {
		if err := s.http.Start(addr); err != nil {
			s.logger.WithError(err).Info("Shutting down the server now")
		}
	}()
}

func (s *Server) watchStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	s.logger.Println(<-stop)
	s.stop()
}

func (s *Server) stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s.logger.Info("Server is stoping...")
	s.http.Shutdown(ctx)
	close(s.signal)
}
