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
	data, err := entity.BindToJsonMap(user)
	if err != nil {
		return err
	}

	doc := r.Client.Collection("users").Doc(user.UserID)

	exist, err := r.checkIfDataExists(doc)
	if err != nil {
		return err
	}
	if exist {
		return ErrUserAlreadyExists
	}

	_, err = doc.Set(r.Context, data)
	return err
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
