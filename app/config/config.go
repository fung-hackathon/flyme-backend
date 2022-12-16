package config

import (
	"fmt"
	"os"
)

var (
	PORT                           string
	PROJECT_ID                     string
	BUCKET_ID                      string
	GOOGLE_APPLICATION_CREDENTIALS string
	YOLP_APPID                     string
	MODE                           Mode
)

type (
	Mode string
)

var (
	Production Mode = "production"
	Developing Mode = "developing"
)

func init() {
	var err error

	PORT, err = getPORT()
	if err != nil {
		panic(err)
	}

	PROJECT_ID, err = getPROJECT_ID()
	if err != nil {
		panic(err)
	}

	BUCKET_ID, err = getBUCKET_ID()
	if err != nil {
		panic(err)
	}

	GOOGLE_APPLICATION_CREDENTIALS, err = getGOOGLE_APPLICATION_CREDENTIALS()
	if err != nil {
		panic(err)
	}

	YOLP_APPID, err = getYOLP_APPID()
	if err != nil {
		panic(err)
	}

	MODE, err = getMODE()
	if err != nil {
		panic(err)
	}
}

func getPORT() (string, error) {
	key := "PORT"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func getPROJECT_ID() (string, error) {
	key := "PROJECT_ID"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func getBUCKET_ID() (string, error) {
	key := "BUCKET_ID"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func getGOOGLE_APPLICATION_CREDENTIALS() (string, error) {
	key := "GOOGLE_APPLICATION_CREDENTIALS"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func getYOLP_APPID() (string, error) {
	key := "YOLP_APPID"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func getMODE() (Mode, error) {
	var m Mode
	if s := os.Getenv("MODE"); s == "production" {
		m = Production
	} else {
		m = Developing
	}
	return m, nil
}
