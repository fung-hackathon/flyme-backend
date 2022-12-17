package usecase

import (
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/domain/repository"
)

type FollowUseCase struct {
	dbRepository repository.DBRepositoryImpl
}

func NewFollowUseCase(r repository.DBRepositoryImpl) *FollowUseCase {
	return &FollowUseCase{
		dbRepository: r,
	}
}

func (u *FollowUseCase) ListFollower(userID string) ([]*entity.GetUser, error) {
	followers, err := u.dbRepository.GetFollowers(userID)
	if err != nil {
		return nil, err
	}

	//friends := make([]response.UserInfo, len(followers.Followers))
	users := make([]*entity.GetUser, len(followers.Followers))

	for i, uid := range followers.Followers {
		user, err := u.dbRepository.GetUser(uid)

		if err != nil {
			return nil, err
		}

		users[i] = user
	}

	return users, nil
}

func (u *FollowUseCase) SendFollow(followeeUserID, followerUserID string) (*entity.GetUser, error) {

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

	return user, nil
}
