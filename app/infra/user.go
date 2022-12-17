package infra

import (
	"flyme-backend/app/domain/entity"

	"cloud.google.com/go/firestore"
)

func (r *DBRepository) GetUser(userID string) (*entity.GetUser, error) {

	doc := r.Client.Collection("users").Doc(userID)

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

	var user entity.GetUser
	err = entity.BindToJsonStruct(docSnap.Data(), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *DBRepository) InsertUser(user *entity.InsertUser) error {

	// user欄を追加
	{
		userDoc := r.Client.Collection("users").Doc(user.UserID)

		exist, err := r.checkIfDataExists(userDoc)
		if err != nil {
			return err
		}
		if exist {
			return ErrUserAlreadyExists
		}

		userTable := entity.UserTable{
			UserID:              user.UserID,
			UserName:            user.UserName,
			Passwd:              user.Passwd,
			Icon:                user.Icon,
			HistoryIDInProgress: "",
		}

		data, err := entity.BindToJsonMap(&userTable)
		if err != nil {
			return err
		}

		_, err = userDoc.Set(r.Context, data)
		if err != nil {
			return err
		}
	}

	// follower欄を追加
	{
		followerDoc := r.Client.Collection("followers").Doc(user.UserID)

		followerTable := entity.FollowerTable{
			Followers: make([]string, 0),
		}

		data, err := entity.BindToJsonMap(&followerTable)
		if err != nil {
			return err
		}

		_, err = followerDoc.Set(r.Context, data)
		if err != nil {
			return err
		}
	}

	// timeline欄を追加
	{
		timelineDoc := r.Client.Collection("timelines").Doc(user.UserID)

		timelineTable := entity.TimelineTable{
			Histories: make([]string, 0),
		}

		data, err := entity.BindToJsonMap(&timelineTable)
		if err != nil {
			return err
		}

		_, err = timelineDoc.Set(r.Context, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *DBRepository) PutUser(user *entity.PutUser) error {

	doc := r.Client.Collection("users").Doc(user.UserID)

	exist, err := r.checkIfDataExists(doc)
	if err != nil {
		return err
	}
	if !exist {
		return ErrUserNotFound
	}

	info := []firestore.Update{
		{Path: "userName", Value: user.UserName},
		{Path: "icon", Value: user.Icon},
	}

	_, err = doc.Update(r.Context, info)
	return err
}
