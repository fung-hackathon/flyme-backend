package interfaces

import (
	"flyme-backend/app/config"
	"flyme-backend/app/infra"
	"flyme-backend/app/interfaces/handler"
	"flyme-backend/app/interfaces/middleware"
	"flyme-backend/app/logger"
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

	userUseCase := usecase.NewUserUseCase(dbRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	followUseCase := usecase.NewFollowUseCase(dbRepository)
	followHandler := handler.NewFollowHandler(followUseCase)

	historyUseCase := usecase.NewHistoryUseCase(dbRepository)
	historyHandler := handler.NewHistoryHandler(historyUseCase)

	bucketRepository, err := infra.NewBucket(ctx, app)
	if err != nil {
		logger.Log{
			Message: "cannot establish DB repository",
			Cause:   err,
		}.Err()
		return
	}

	imageUsecase := usecase.NewImageUseCase(bucketRepository, dbRepository)
	imgHandler := handler.NewImageHandler(imageUsecase)

	s.Router.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	s.Router.POST("/user", userHandler.CreateUser)
	s.Router.POST("/login", userHandler.Login)
	s.Router.GET("/user/:user_id", userHandler.ReadUser)

	ur := s.Router.Group("")
	{
		ur.Use(middleware.Authentication("user"))
		ur.PUT("/user/:user_id", userHandler.UpdateUser)
		ur.POST("/icon/:user_id", imgHandler.UploadIcon)
		ur.GET("/icon/:user_id", imgHandler.DownloadIcon)
	}

	fr := s.Router.Group("/follow")
	{
		fr.Use(middleware.Authentication("user"))

		fr.GET("/:user_id", followHandler.ListFollower)
		fr.POST("/:user_id", followHandler.SendFollow)
	}

	hr := s.Router.Group("/history")
	{
		hr.Use(middleware.Authentication("user"))

		hr.POST("/:user_id/start", historyHandler.StartHistory)
		hr.POST("/:user_id/finish", historyHandler.FinishHistory)
		hr.GET("/:user_id", historyHandler.ReadHistories)
		hr.GET("/:user_id/timeline", historyHandler.ReadTimeline)
	}

	if config.MODE == config.Production {
		s.Router.HideBanner = true
		s.Router.HidePort = true
	}

	s.Router.Start(":" + config.PORT)
}
