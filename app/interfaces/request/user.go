package request

type CreateUserRequest struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
}

type UpdateUserRequest struct {
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}
