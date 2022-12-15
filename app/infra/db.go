package infra

import (
	"context"
	"errors"
	"flyme-backend/app/config"
	"flyme-backend/app/domain/entity"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrFirebaseInitializaition = errors.New("failed to initialize Firebase")
	ErrFirestoreConnection     = errors.New("failed to establish connection to Firestore")
	ErrUserNotFound            = errors.New("user not found")
	ErrUserAlreadyExists       = errors.New("user already exists")
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

func (r *DBRepository) checkIfDataExists(doc *firestore.DocumentRef) (bool, error) {
	_, err := doc.Get(r.Context)
	if err == nil {
		return true, nil
	} else if status.Code(err) == codes.NotFound {
		return false, nil
	} else {
		return false, err
	}
}

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
