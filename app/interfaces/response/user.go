package response

type UserInfo struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}

type ReadUserResponse = UserInfo
type CreateUserResponse = UserInfo
type UpdateUserResponse = UserInfo

type LoginResponse struct {
	Token string `json:"token"`
}

type ValidateUserTokenResponse struct {
	Valid bool `json:"valid"`
}
