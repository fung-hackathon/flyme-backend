package handler

import (
	"flyme-backend/app/interfaces/request"
	"flyme-backend/app/interfaces/response"
	"flyme-backend/app/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HistoryHandler struct {
	historyUseCase *usecase.HistoryUseCase
}

func NewHistoryHandler(u *usecase.HistoryUseCase) *HistoryHandler {
	return &HistoryHandler{
		historyUseCase: u,
	}
}

func (h *HistoryHandler) StartHistory(c echo.Context) error {
	userID := c.Param("user_id")

	var req request.StartHistoryRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	history, err := h.historyUseCase.StartHistory(userID, &req)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, history)
}

func (h *HistoryHandler) FinishHistory(c echo.Context) error {
	userID := c.Param("user_id")

	var req request.FinishHistoryRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	history, err := h.historyUseCase.FinishHistory(userID, &req)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, history)
}

func (h *HistoryHandler) ReadHistories(c echo.Context) error {
	userID := c.Param("user_id")
	size, err := strconv.ParseInt(c.QueryParam("number"), 10, 32)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	histories, err := h.historyUseCase.ReadHistories(userID, int(size))
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, histories)
}

func (h *HistoryHandler) ReadTimeline(c echo.Context) error {
	userID := c.Param("user_id")
	size, err := strconv.ParseInt(c.QueryParam("number"), 10, 32)

	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			response.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		)
	}

	timeline, err := h.historyUseCase.ReadTimeline(userID, int(size))
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, timeline)
}
