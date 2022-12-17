package entity

type TimelineTable struct {
	Histories []string `json:"histories"`
}

type GetTimeline = TimelineTable
