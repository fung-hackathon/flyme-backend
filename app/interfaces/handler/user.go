package handler

import (
	"net/http"

	"flyme-backend/app/interfaces/request"
	"flyme-backend/app/interfaces/response"
	"flyme-backend/app/logger"
	"flyme-backend/app/packages/auth"
	"flyme-backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: u,
	}
}

func (h *UserHandler) ReadUser(c echo.Context) error {
	userID := c.Param("user_id")
	user, err := h.userUseCase.ReadUser(userID)
	if err != nil {
		logger.Log{
			Message: "read user was failed",
			Cause:   err,
		}.Warn()
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	response := &response.ReadUserResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c echo.Context) error {

	var req request.CreateUserRequest

	if err := c.Bind(&req); err != nil {
		logger.Log{
			Message: "unexpected request body",
			Cause:   err,
		}.Warn()
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	query, err := h.userUseCase.CreateUser(&req)
	if err != nil {
		logger.Log{
			Message: "create user was failed",
			Cause:   err,
		}.Warn()
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	response := &response.CreateUserResponse{
		UserID:   query.UserID,
		UserName: query.UserName,
		Icon:     query.Icon,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {

	userID := c.Param("user_id")

	claims, err := auth.GetUserContext(c.Get("user"))

	if err != nil {
		logger.Log{
			Message: "auth was failed",
			Cause:   err,
		}.Warn()
		return c.JSON(
			http.StatusUnauthorized,
			response.Error{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			},
		)
	}

	if claims.UserID != userID {
		return c.JSON(
			http.StatusUnauthorized,
			response.Error{
				Code:    http.StatusUnauthorized,
				Message: "not authorized",
			},
		)
	}

	var req request.UpdateUserRequest

	if err := c.Bind(&req); err != nil {
		logger.Log{
			Message: "unexpected request body",
			Cause:   err,
		}.Warn()
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	user, err := h.userUseCase.UpdateUser(userID, &req)
	if err != nil {
		logger.Log{
			Message: "update user was failed",
			Cause:   err,
		}.Err()
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	response := &response.ReadUserResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) Login(c echo.Context) error {

	var req request.LoginRequest

	if err := c.Bind(&req); err != nil {
		logger.Log{
			Message: "unexpected request body",
			Cause:   err,
		}.Warn()
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	logger.RequestLog(req)

	token, err := h.userUseCase.Login(&req)
	if err != nil {
		logger.Log{
			Message: "unexpected request body",
			Cause:   err,
		}.Warn()
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	response := &response.LoginResponse{
		Token: token,
	}

	return c.JSON(http.StatusOK, response)
}
