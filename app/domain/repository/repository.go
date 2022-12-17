package repository

import (
	"flyme-backend/app/domain/entity"
	"io"
)

type DBRepositoryImpl interface {
	GetUser(string) (*entity.GetUser, error)
	InsertUser(*entity.InsertUser) error
	PutUser(*entity.PutUser) error

	GetFollowers(string) (*entity.GetFollowers, error)
	SendFollow(*entity.SendFollow) error

	GetHistory(string) (*entity.GetHistory, error)
	StartHistory(*entity.StartHistory) (*entity.HistoryTable, error)
	FinishHistory(*entity.FinishHistory) (*entity.HistoryTable, error)

	GetHistories(string, int) (*entity.GetHistories, error)
	GetTimeline(string, int) (*entity.GetTimeline, error)

	CheckUserExist(userID string) (bool, error)
}

type ImageRepositoryImpl interface {
	UploadIconImg(file io.Reader, userID string) error
	DownloadIconImg(file io.Writer, userID string) error
}
