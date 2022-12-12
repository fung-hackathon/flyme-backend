package infra

import (
	"context"
	"errors"
	"flyme-backend/app/config"
	"flyme-backend/app/domain/entity"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	ErrFirebaseInitializaition = errors.New("failed to initialize Firebase")
	ErrFirestoreConnection     = errors.New("failed to establish connection to Firestore")
)

type DBRepository struct {
	Client  *firestore.Client
	Context context.Context
}

func NewDBRepository() (*DBRepository, error) {
	opt := option.WithCredentialsFile(config.GOOGLE_APPLICATION_CREDENTIALS)
	conf := &firebase.Config{ProjectID: config.PROJECT_ID}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, ErrFirebaseInitializaition
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, ErrFirestoreConnection
	}

	return &DBRepository{Client: client, Context: ctx}, nil
}

func (r *DBRepository) Close() {
	r.Client.Close()
}

func (r *DBRepository) GetUser(userID string) (*entity.GetUser, error) {
	dsnap, err := r.Client.Collection("users").Doc(userID).Get(r.Context)
	if err != nil {
		return nil, err
	}

	var user entity.GetUser
	err = entity.BindToJson(dsnap.Data(), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *DBRepository) InsertUser(user *entity.InsertUser) error {
	_, err := r.Client.Collection("users").Doc(user.UserID).Set(r.Context, user)
	return err
}

func (r *DBRepository) PutUser(user *entity.PutUser) error {
	info := []firestore.Update{{
		Path:  "*",
		Value: user,
	}}
	_, err := r.Client.Collection("users").Doc(user.UserID).Update(r.Context, info)
	return err
}
