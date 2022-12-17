package auth

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"flyme-backend/app/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserJwtClaims struct {
	UserID string `json:"userID"`
}

var signingKey []byte

func init() {
	if config.MODE == config.Production {
		signingKey = make([]byte, 128)
		_, err := rand.Read(signingKey)
		if err != nil {
			panic(err)
		}
	} else {
		signingKey = []byte("nannde")
	}
}

func GenerateUserToken(userID string, passwd string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenStr, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateUserToken(tokenStr string) (interface{}, error) {

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, errors.New("unexpected jwt signing method")
		}
		return signingKey, nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func GetUserContext(src interface{}) (*UserJwtClaims, error) {
	if src == nil {
		return nil, errors.New("failed to resolve user context")
	}
	token := src.(*jwt.Token)

	jsonStr, err := json.Marshal(token.Claims)
	if err != nil {
		return nil, err
	}

	var claims UserJwtClaims
	err = json.Unmarshal(jsonStr, &claims)
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
