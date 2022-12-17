package request

type LoginRequest struct {
	UserID string `json:"userID"`
	Passwd string `json:"passwd"`
}

type CreateUserRequest struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
}

type UpdateUserRequest struct {
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}
