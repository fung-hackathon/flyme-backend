package entity

type Coordinate struct {
	Longitude float64 `json:"lng"`
	Latitude  float64 `json:"lat"`
}

type GetHistory struct {
	Coords    []Coordinate `json:"coordinates"`
	Dist      string       `json:"dist"`
	Finish    string       `json:"finish"`
	Start     string       `json:"start"`
	State     string       `json:"state"`
	UserID    string       `json:"userID"`
	HistoryID string       `json:"historyID"`
}

type InsertHistory struct {
	Coords    []Coordinate `json:"coordinates"`
	State     string       `json:"state"`
	UserID    string       `json:"userID"`
	HistoryID string       `json:"historyID"`
}
