package infra

import (
	"context"
	"flyme-backend/app/config"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func FirebaseNewApp() (context.Context, *firebase.App, error) {
	opt := option.WithCredentialsFile(config.GOOGLE_APPLICATION_CREDENTIALS)
	conf := &firebase.Config{
		ProjectID:     config.PROJECT_ID,
		StorageBucket: config.BUCKET_ID,
	}

	ctx := context.Background()
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, nil, err
	}
	return ctx, app, nil
}
