package response

type ReadUserResponse struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}

type CreateUserResponse struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}

type UpdateUserResponse struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}
