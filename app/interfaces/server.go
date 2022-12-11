package interfaces

import (
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
	s.Router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	s.Router.Start(":3000")
}
