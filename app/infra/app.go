package infra

import (
	"context"
	"flyme-backend/app/config"
	"flyme-backend/app/logger"

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
		logger.Log{
			Message: "failed to create Firebase App",
			Cause:   err,
		}.Err()
		return nil, nil, err
	}

	logger.Log{
		Message: "created Firebase App",
	}.Info()
	return ctx, app, nil
}
