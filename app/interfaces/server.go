package interfaces

import (
	"flyme-backend/app/config"
	"flyme-backend/app/infra"
	"flyme-backend/app/interfaces/handler"
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

	dbRepository, err := infra.NewDBRepository()
	if err != nil {
		s.Router.Logger.Error(err)
		return
	}

	userUseCase := usecase.NewUseCase(dbRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	s.Router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	s.Router.GET("/user/:user_id", userHandler.ReadUser)
	s.Router.POST("/user", userHandler.CreateUser)

	s.Router.Logger.Fatal(s.Router.Start(":" + config.PORT))

}
