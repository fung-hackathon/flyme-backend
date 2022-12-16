package entity

type GetFollowers struct {
	Followers []string `json:"followers"`
}

type SendFollow struct {
	FolloweeUserID string `json:"followeeUserID"`
	FollowerUserID string `json:"followerUserID"`
}
