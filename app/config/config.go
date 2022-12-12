package config

import (
	"os"
)

var (
	PORT                           string
	PROJECT_ID                     string
	GOOGLE_APPLICATION_CREDENTIALS string
)

func init() {
	PORT = os.Getenv("PORT")
	PROJECT_ID = os.Getenv("PROJECT_ID")
	GOOGLE_APPLICATION_CREDENTIALS = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
}
