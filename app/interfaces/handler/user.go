package handler

import (
	"net/http"

	"flyme-backend/app/interfaces/request"
	"flyme-backend/app/interfaces/response"
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
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	user, err := h.userUseCase.CreateUser(&req)
	if err != nil {
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

func (h *UserHandler) UpdateUser(c echo.Context) error {

	userID := c.Param("user_id")

	var req request.UpdateUserRequest

	if err := c.Bind(&req); err != nil {
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
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	info, err := h.userUseCase.Login(&req)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	response := &response.LoginResponse{
		Token: info.Token,
	}

	return c.JSON(http.StatusOK, response)
}
