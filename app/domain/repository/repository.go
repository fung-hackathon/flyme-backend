package repository

import "flyme-backend/app/domain/entity"

type DBRepositoryImpl interface {
	GetUser(string) (*entity.GetUser, error)
	InsertUser(*entity.InsertUser) error
	PutUser(*entity.PutUser) error

	GetFollowers(string) (*entity.GetFollowers, error)
	SendFollow(*entity.SendFollow) error

	GetHistory(string) (*entity.GetHistory, error)
	InsertHistory(*entity.InsertHistory) (*entity.GetHistory, error)

	GetTimeline(string) (*entity.GetTimeline, error)
}
