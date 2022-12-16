package entity

type UserTable struct {
	UserID              string `json:"userID"`
	UserName            string `json:"userName"`
	Passwd              string `json:"passwd"`
	HistoryIDInProgress string `json:"historyIDInProgress"`
	Icon                string `json:"icon"`
}

type GetUser = UserTable

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
