package handler

import (
	"flyme-backend/app/interfaces/response"
	"flyme-backend/app/packages/auth"
	"flyme-backend/app/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FollowHandler struct {
	followUseCase *usecase.FollowUseCase
}

func NewFollowHandler(u *usecase.FollowUseCase) *FollowHandler {
	return &FollowHandler{
		followUseCase: u,
	}
}

func (h *FollowHandler) ListFollower(c echo.Context) error {

	userID := c.Param("user_id")
	friends, err := h.followUseCase.ListFollower(userID)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, friends)
}

func (h *FollowHandler) SendFollow(c echo.Context) error {

	followerUserID := c.Param("user_id")

	claims, err := auth.GetUserContext(c.Get("user"))

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	followeeUserID := claims.UserID

	user, err := h.followUseCase.SendFollow(followeeUserID, followerUserID)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, user)
}
