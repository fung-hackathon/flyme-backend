package infra

import (
	"flyme-backend/app/domain/entity"
)

func (r *DBRepository) GetFollowers(userID string) (*entity.GetFollowers, error) {
	return nil, nil
}

func (r *DBRepository) SendFollow(follow *entity.SendFollow) error {
	return nil
}
