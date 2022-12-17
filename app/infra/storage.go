package infra

import (
	"context"
	"errors"
	"flyme-backend/app/logger"
	"io"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
)

var (
	ErrFirebaseStorageInit   = errors.New("failed to initialize Firebase Storage")
	ErrFirebaseStorageBucket = errors.New("failed to initialize Firebase Storage Bucket")
)

type BucketRepository struct {
	bucket *storage.BucketHandle
}

func NewBucket(ctx context.Context, app *firebase.App) (*BucketRepository, error) {
	client, err := app.Storage(context.Background())
	if err != nil {
		logger.Log{
			Message: "failed to create new Firestorage client",
			Cause:   err,
		}.Err()
		return nil, err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		logger.Log{
			Message: "failed to create new Bucket",
			Cause:   err,
		}.Err()
		return nil, err
	}

	logger.Log{
		Message: "created new Bucket",
	}.Info()

	return &BucketRepository{bucket}, nil
}

func (r *BucketRepository) UploadIconImg(file io.Reader, userID string) error {
	contentType := "image/png"
	ctx := context.Background()
	writer := r.bucket.Object("icon/" + userID + ".png").NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}

	if _, err := io.Copy(writer, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	logger.Log{
		Message: "uploaded new icon to Firestorage",
	}.Info()

	return nil
}

func (r *BucketRepository) DownloadIconImg(file io.Writer, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	rc, err := r.bucket.Object("icon/" + userID + ".png").NewReader(ctx)
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, rc); err != nil {
		return err
	}

	logger.Log{
		Message: "downloaded new icon from Firestorage",
	}.Info()

	return nil
}
