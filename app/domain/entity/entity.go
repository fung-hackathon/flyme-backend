package entity

type User struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
	Icon     string `json:"icon"`
}
