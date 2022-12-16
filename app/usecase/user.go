package usecase

import (
	"errors"
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/interfaces/request"
	"flyme-backend/app/interfaces/response"
	"flyme-backend/app/packages/auth"
)

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

func (u *UserUseCase) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	user, err := u.dbRepository.GetUser(req.UserID)

	if err != nil {
		return nil, err
	}

	if req.Passwd != user.Passwd {
		return nil, errors.New("password incorrect")
	}

	token, err := auth.GenerateUserToken(req.UserID, req.Passwd)
	if err != nil {
		return nil, err
	}

	res := &response.LoginResponse{
		Token: token,
	}

	return res, nil
}
