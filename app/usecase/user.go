package usecase

import (
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/domain/repository"
	"flyme-backend/app/interfaces/request"
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

func (u *UserUseCase) ReadUser(userID string) (*response.ReadUserResponse, error) {
	user, err := u.dbRepository.GetUser(userID)

	if err != nil {
		return nil, err
	}

	res := &response.ReadUserResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		Icon:     user.Icon,
	}

	return res, nil
}

func (u *UserUseCase) CreateUser(req *request.CreateUserRequest) (*response.CreateUserResponse, error) {

	// TODO: Default Icon

	query := &entity.InsertUser{
		UserID:   req.UserID,
		UserName: req.UserName,
		Passwd:   req.Passwd,
		Icon:     "",
	}

	err := u.dbRepository.InsertUser(query)
	if err != nil {
		return nil, err
	}

	res := &response.CreateUserResponse{
		UserID:   query.UserID,
		UserName: query.UserName,
		Icon:     query.Icon,
	}

	return res, nil
}

func (u *UserUseCase) UpdateUser(userID string, req *request.UpdateUserRequest) (*response.UpdateUserResponse, error) {

	query := &entity.PutUser{
		UserID:   userID,
		UserName: req.UserName,
		Icon:     req.Icon,
	}

	err := u.dbRepository.PutUser(query)
	if err != nil {
		return nil, err
	}

	res := &response.UpdateUserResponse{
		UserID:   query.UserID,
		UserName: query.UserName,
		Icon:     query.Icon,
	}

	return res, nil
}
