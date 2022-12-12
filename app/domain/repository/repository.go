package repository

import "flyme-backend/app/domain/entity"

type DBRepositoryImpl interface {
	GetUser(string) (*entity.GetUser, error)
}
