package repository

import "flyme-backend/app/domain/entity"

type DBRepositoryImpl interface {
	GetUser(string) (*entity.GetUser, error)
	InsertUser(*entity.InsertUser) error
	PutUser(*entity.PutUser) error
}
