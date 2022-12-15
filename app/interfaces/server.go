package interfaces

import (
	"flyme-backend/app/config"
	"flyme-backend/app/infra"
	"flyme-backend/app/interfaces/handler"
	"flyme-backend/app/interfaces/middleware"
	"flyme-backend/app/logger"
	"flyme-backend/app/packages/auth"
	"flyme-backend/app/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Router *echo.Echo
}

func NewServer() *Server {
	return &Server{echo.New()}
}

func (s *Server) StartServer() {
	s.Router.Use(logger.EchoLogger())

	if config.MODE == config.Developing {
		s.Router.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		}))
	}

	ctx, app, err := infra.FirebaseNewApp()

	if err != nil {
		logger.Log{
			Message: "failed to initialize Firebase",
			Cause:   err,
		}.Err()
		return
	}

	dbRepository, err := infra.NewDBRepository(ctx, app)

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

	s.Router.POST("/user", userHandler.CreateUser)
	s.Router.POST("/login", userHandler.Login)
	s.Router.GET("/user/:user_id", userHandler.ReadUser)
	s.Router.PUT("/user/:user_id", userHandler.UpdateUser)

	// authorized '/ping' ---
	r := s.Router.Group("/auth")
	{
		const contextKey = "user"
		r.Use(middleware.Authentication(contextKey))

		r.GET("/ping", func(c echo.Context) error {
			ctx, _ := auth.GetUserContext(c.Get(contextKey))
			return c.String(http.StatusOK, "pong by "+ctx.UserID)
		})
	}
	// ---

	if config.MODE == config.Production {
		s.Router.HideBanner = true
		s.Router.HidePort = true
	}

	s.Router.Start(":" + config.PORT)
}
