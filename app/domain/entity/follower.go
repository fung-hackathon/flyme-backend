package entity

type GetFollowers struct {
	UserID []string `json:"followers"`
}

type SendFollow struct {
	FolloweeUserID string `json:"followeeUserID"`
	FollowerUserID string `json:"followerUserID"`
}
