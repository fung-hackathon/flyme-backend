package interfaces

import (
	"flyme-backend/app/config"
	"flyme-backend/app/infra"
	"flyme-backend/app/interfaces/handler"
	"flyme-backend/app/logger"
	"flyme-backend/app/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func NewServer() *Server {
	return &Server{echo.New()}
}

func (s *Server) StartServer() {
	s.Router.Use(logger.EchoLogger())

	dbRepository, err := infra.NewDBRepository()
	if err != nil {
		logger.Log{
			Message: "cannot establish DB repository",
			Cause:   err,
		}.Err()
		return
	}

	userUseCase := usecase.NewUseCase(dbRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	s.Router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	s.Router.GET("/user/:user_id", userHandler.ReadUser)
	s.Router.PUT("/user/:user_id", userHandler.UpdateUser)
	s.Router.POST("/user", userHandler.CreateUser)

	if config.MODE == config.Production {
		s.Router.HideBanner = true
		s.Router.HidePort = true
	}

	s.Router.Start(":" + config.PORT)
}
