package usecase

import (
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/domain/repository"
	"flyme-backend/app/interfaces/response"
)

type FollowUseCase struct {
	dbRepository repository.DBRepositoryImpl
}

func NewFollowUseCase(r repository.DBRepositoryImpl) *FollowUseCase {
	return &FollowUseCase{
		dbRepository: r,
	}
}

func (u *FollowUseCase) ListFollower(userID string) (*response.ListFollowerResponse, error) {
	followers, err := u.dbRepository.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	friends := make([]response.UserInfo, len(followers.Followers))

	for i, uid := range followers.Followers {
		user, err := u.dbRepository.GetUser(uid)

		if err != nil {
			return nil, err
		}

		friends[i] = response.UserInfo{
			UserID:   user.UserID,
			UserName: user.UserName,
			Icon:     user.Icon,
		}
	}

	res := &response.ListFollowerResponse{
		Friends: friends,
	}

	return res, nil
}

func (u *FollowUseCase) SendFollow(followeeUserID, followerUserID string) (*response.SendFollowResponse, error) {

	err := u.dbRepository.SendFollow(&entity.SendFollow{
		FolloweeUserID: followeeUserID,
		FollowerUserID: followerUserID,
	})

	if err != nil {
		return nil, err
	}

	user, err := u.dbRepository.GetUser(followerUserID)

	if err != nil {
		return nil, err
	}

	res := &response.SendFollowResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}

	return res, nil
}
