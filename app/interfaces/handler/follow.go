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

	claims, err := auth.GetUserContext(c.Get("user"))

	if err != nil {
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

	followers, err := h.followUseCase.ListFollower(userID)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	friends := make([]response.UserInfo, len(followers))

	for i, user := range followers {
		friends[i] = response.UserInfo{
			UserID:   user.UserID,
			UserName: user.UserName,
			Icon:     user.Icon,
		}
	}

	response := &response.ListFollowerResponse{
		Friends: friends,
	}

	return c.JSON(http.StatusOK, response)
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

	if followeeUserID == followerUserID {
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: "cannot follow yourself",
			},
		)
	}

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

	response := &response.SendFollowResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}

	return c.JSON(http.StatusOK, response)
}
