package infra

import (
	"context"
	"flyme-backend/app/logger"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DBRepository struct {
	Client  *firestore.Client
	Context context.Context
}

func NewDBRepository(ctx context.Context, app *firebase.App) (*DBRepository, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		logger.Log{
			Message: "failed to create connection to Firestore",
			Cause:   err,
		}.Err()
		return nil, ErrFirestoreConnection
	}
	logger.Log{
		Message: "created connection to Firestore",
	}.Info()
	return &DBRepository{Client: client, Context: ctx}, nil
}

func (r *DBRepository) Close() {
	r.Client.Close()
}

func (r *DBRepository) CheckUserExist(userID string) (bool, error) {
	doc := r.Client.Collection("users").Doc(userID)

	_, err := doc.Get(r.Context)
	if err == nil {
		return true, nil
	} else if status.Code(err) == codes.NotFound {
		return false, nil
	} else {
		return false, err
	}
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
