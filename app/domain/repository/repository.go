package repository

import (
	"flyme-backend/app/domain/entity"
	"io"
)

type DBRepositoryImpl interface {
	GetUser(string) (*entity.GetUser, error)
	InsertUser(*entity.InsertUser) error
	PutUser(*entity.PutUser) error
	CheckUserExist(userID string) (bool, error)
}

type ImageRepositoryImpl interface {
	UploadIconImg(file io.Reader, userID string) error
	DownloadIconImg(file io.Writer, userID string) error
}
