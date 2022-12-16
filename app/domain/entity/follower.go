package entity

type FollowerTable struct {
	Followers []string `json:"followers"`
}

type GetFollowers = FollowerTable

type SendFollow struct {
	FolloweeUserID string `json:"followeeUserID"`
	FollowerUserID string `json:"followerUserID"`
}
