package request

type Coordinate struct {
	Longitude float64 `json:"lng"`
	Latitude  float64 `json:"lat"`
}

type StartHistoryRequest struct {
	StartTime string `json:"start"`
	Ticket    string `json:"ticket"`
}

type FinishHistoryRequest struct {
	Coords     []Coordinate `json:"coordinates"`
	Distance   float64      `json:"dist"`
	FinishTime string       `json:"finish"`
}
