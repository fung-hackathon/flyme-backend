package usecase

import (
	"flyme-backend/app/domain/repository"
)

type UserUseCase struct {
	dbRepository repository.DBRepositoryImpl
}

func NewUseCase(r repository.DBRepositoryImpl) *UserUseCase {
	return &UserUseCase{
		dbRepository: r,
	}
}
