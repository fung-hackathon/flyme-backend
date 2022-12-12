package config

import (
	"fmt"
	"os"
)

var (
	PORT                           string
	PROJECT_ID                     string
	GOOGLE_APPLICATION_CREDENTIALS string
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

	PORT, err = GetPORT()
	if err != nil {
		panic(err)
	}

	PROJECT_ID, err = GetPROJECT_ID()
	if err != nil {
		panic(err)

	}

	GOOGLE_APPLICATION_CREDENTIALS, err = GetGOOGLE_APPLICATION_CREDENTIALS()
	if err != nil {
		panic(err)
	}

	MODE, err = GetMODE()
	if err != nil {
		panic(err)

	}
}

func GetPORT() (string, error) {
	key := "PORT"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func GetPROJECT_ID() (string, error) {
	key := "PROJECT_ID"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func GetGOOGLE_APPLICATION_CREDENTIALS() (string, error) {
	key := "GOOGLE_APPLICATION_CREDENTIALS"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func GetMODE() (Mode, error) {
	var m Mode
	if s := os.Getenv("MODE"); s == "production" {
		m = Production
	} else {
		m = Developing
	}
	return m, nil
}
