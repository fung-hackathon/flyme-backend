package handler

import (
	"flyme-backend/app/interfaces/request"
	"flyme-backend/app/interfaces/response"
	"flyme-backend/app/packages/auth"
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
			http.StatusInternalServerError,
			response.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	rcoords := make([]response.Coordinate, len(history.Coords))

	for i, c := range history.Coords {
		rcoords[i] = response.Coordinate{
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
		}
	}

	response := &response.StartHistoryResponse{
		Coords: rcoords,
		Dist:   history.Dist,
		Finish: history.Finish,
		Start:  history.Start,
		State:  history.State,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *HistoryHandler) FinishHistory(c echo.Context) error {
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
			http.StatusInternalServerError,
			response.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		)
	}

	rcoords := make([]response.Coordinate, len(req.Coords))

	for i, c := range history.Coords {
		rcoords[i] = response.Coordinate{
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
		}
	}

	response := &response.FinishHistoryResponse{
		Coords: rcoords,
		Dist:   history.Dist,
		Finish: history.Finish,
		Start:  history.Start,
		State:  history.State,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *HistoryHandler) ReadHistories(c echo.Context) error {

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

	rtimeline := []response.HistoryTable{}

	for _, history := range histories.Histories {

		rcoords := make([]response.Coordinate, len(history.Coords))

		for i, c := range history.Coords {
			rcoords[i] = response.Coordinate{
				Longitude: c.Longitude,
				Latitude:  c.Latitude,
			}
		}

		rtimeline = append(rtimeline, response.HistoryTable{
			Coords: rcoords,
			Dist:   history.Dist,
			Finish: history.Finish,
			Start:  history.Start,
			State:  history.Start,
		})
	}

	response := &response.ReadHistoriesResponse{
		Histories: rtimeline,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *HistoryHandler) ReadTimeline(c echo.Context) error {
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

	histories, users, err := h.historyUseCase.ReadTimeline(userID, int(size))
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			response.Error{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			},
		)
	}

	timeline := []response.HistoryTimeline{}

	for i := range timeline {

		timeline = append(timeline, response.HistoryTimeline{
			User: response.UserInfo{
				UserID:   users[i].UserID,
				UserName: users[i].UserName,
				Icon:     users[i].Icon,
			},
			Finish: histories[i].Finish,
			Start:  histories[i].Start,
			State:  histories[i].State,
		})
	}

	response := &response.ReadTimelineResponse{
		Histories: timeline,
	}

	return c.JSON(http.StatusOK, response)
}
