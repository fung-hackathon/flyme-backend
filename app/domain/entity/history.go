package entity

import "github.com/google/uuid"

type Coordinate struct {
	Longitude float64 `json:"Longitude"`
	Latitude  float64 `json:"Latitude"`
}

type HistoryTable struct {
	Coords    []Coordinate `json:"coordinates"`
	Dist      float64      `json:"dist"`
	Finish    string       `json:"finish"`
	Start     string       `json:"start"`
	State     string       `json:"state"`
	Ticket    string       `json:"ticket"`
	UserID    string       `json:"userID"`
	HistoryID string       `json:"historyID"`
}

type GetHistory = HistoryTable

type GetHistories struct {
	Histories []GetHistory
}

type StartHistory struct {
	UserID    string
	StartTime string
	Ticket    string
}

type FinishHistory struct {
	Coords     []Coordinate
	UserID     string
	Distance   float64
	FinishTime string
}

func NewHistoryID() string {
	return uuid.New().String()
}
