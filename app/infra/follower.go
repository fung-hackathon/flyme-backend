package infra

import (
	"flyme-backend/app/domain/entity"
)

func (r *DBRepository) GetFollowers(userID string) (*entity.GetFollowers, error) {
	doc := r.Client.Collection("followers").Doc(userID)

	exist, err := r.checkIfDataExists(doc)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, ErrUserNotFound
	}

	docSnap, err := doc.Get(r.Context)
	if err != nil {
		return nil, err
	}

	var followers entity.GetFollowers
	err = entity.BindToJsonStruct(docSnap.Data(), &followers)
	if err != nil {
		return nil, err
	}

	return &followers, nil
}

func (r *DBRepository) SendFollow(follow *entity.SendFollow) error {
	doc := r.Client.Collection("followers").Doc(follow.FollowerUserID)

	exist, err := r.checkIfDataExists(doc)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotFound
	}

	docSnap, err := doc.Get(r.Context)
	if err != nil {
		return err
	}

	var followers entity.FollowerTable
	err = entity.BindToJsonStruct(docSnap.Data(), &followers)
	if err != nil {
		return err
	}

	for _, uid := range followers.Followers {
		if uid == follow.FolloweeUserID {
			return nil
		}
	}

	followers.Followers = append(followers.Followers, follow.FolloweeUserID)

	data, err := entity.BindToJsonMap(&followers)
	if err != nil {
		return err
	}

	_, err = doc.Set(r.Context, data)

	return err
}
