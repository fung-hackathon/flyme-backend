package request

type Coordinate struct {
	Longitude float64 `json:"lng"`
	Latitude  float64 `json:"lat"`
}

type StartHistoryRequest struct {
	StartTime string `json:"start"`
}

type FinishHistoryRequest struct {
	Coords     []Coordinate `json:"coordinates"`
	FinishTime string       `json:"finish"`
}
