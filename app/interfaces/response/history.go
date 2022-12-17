package response

type Coordinate struct {
	Longitude float64 `json:"lng"`
	Latitude  float64 `json:"lat"`
}

type HistoryTable struct {
	Coords []Coordinate `json:"coordinates"`
	Dist   float64      `json:"dist"`
	Finish string       `json:"finish"`
	Start  string       `json:"start"`
	State  string       `json:"state"`
}

type HistoryTimeline struct {
	User   UserInfo `json:"user"`
	Finish string   `json:"finish"`
	Start  string   `json:"start"`
	State  string   `json:"state"`
}

type StartHistoryResponse = HistoryTable
type FinishHistoryResponse = HistoryTable

type ReadTimelineResponse struct {
	Histories []HistoryTimeline `json:"histories"`
}
