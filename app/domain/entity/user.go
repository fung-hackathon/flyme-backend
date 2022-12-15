package entity

type GetUser struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
	Icon     string `json:"icon"`
}

type InsertUser struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
	Icon     string `json:"icon"`
}

type PutUser struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}
