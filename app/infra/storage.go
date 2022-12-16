package infra

import (
	"context"
	"errors"
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

func init() {
	// ctx, app, err := FirebaseNewApp()
	// if err != nil {
	// 	logger.Log{
	// 		Message: "Open file",
	// 		Cause:   err,
	// 	}.Err()
	// 	return
	// }

	// r, err := NewBucket(ctx, app)
	// if err != nil {
	// 	logger.Log{
	// 		Message: "Open file",
	// 		Cause:   err,
	// 	}.Err()
	// 	return
	// }

	// f, err := os.Open("account.png")
	// if err != nil {
	// 	logger.Log{
	// 		Message: "Open file",
	// 		Cause:   err,
	// 	}.Err()
	// 	return
	// }
	// defer f.Close()

	// err = r.InsertIconImg(f, "account.png")
	// if err != nil {
	// 	logger.Log{
	// 		Message: "Open file",
	// 		Cause:   err,
	// 	}.Err()
	// 	return
	// }

	// f, err := os.Create("hoge.png")
	// if err != nil {
	// 	logger.Log{
	// 		Message: "Open file",
	// 		Cause:   err,
	// 	}.Err()
	// 	return
	// }

	// err = r.SelectIconImg(f, "hoge")
	// if err != nil {
	// 	logger.Log{
	// 		Message: "Open file",
	// 		Cause:   err,
	// 	}.Err()
	// 	return
	// }
}

func NewBucket(ctx context.Context, app *firebase.App) (*BucketRepository, error) {
	client, err := app.Storage(context.Background())
	if err != nil {
		return nil, err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, err
	}

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

	return nil
}

func (r *BucketRepository) DownloadIconImg(file io.Writer, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	rc, err := r.bucket.Object("icon/" + userID + ".png").NewReader(ctx)
	if err != nil {
		return err
	}

	io.Copy(file, rc)

	return nil
}