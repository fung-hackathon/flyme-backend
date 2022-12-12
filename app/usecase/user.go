package usecase

import (
	"flyme-backend/app/domain/repository"
	"flyme-backend/app/interfaces/response"
)

type UserUseCase struct {
	dbRepository repository.DBRepositoryImpl
}

func NewUseCase(r repository.DBRepositoryImpl) *UserUseCase {
	return &UserUseCase{
		dbRepository: r,
	}
}

func (u *UserUseCase) ReadUser(userID string) (*response.UserResponse, error) {
	user, err := u.dbRepository.GetUser(userID)

	if err != nil {
		return nil, err
	}

	response := &response.UserResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}

	return response, nil
}
