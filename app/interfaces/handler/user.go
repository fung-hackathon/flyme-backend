package handler

import (
	"net/http"

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

	response := &response.UserResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}

	return c.JSON(http.StatusOK, response)
}
