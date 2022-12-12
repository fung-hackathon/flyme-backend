package infra

import (
	"context"
	"encoding/json"
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

	jsonStr, err := json.Marshal(dsnap.Data())
	if err != nil {
		return nil, err
	}

	user := new(entity.GetUser)

	err = json.Unmarshal(jsonStr, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
