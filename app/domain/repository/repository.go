package repository

import "flyme-backend/app/domain/entity"

type IDBRepository interface {
	GetUser(string) (*entity.User, error)
}
