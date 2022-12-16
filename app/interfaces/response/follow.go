package response

type ListFollowerResponse struct {
	Friends []UserInfo `json:"friends"`
}

type SendFollowResponse = UserInfo
