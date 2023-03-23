package usecase

import (
	"errors"
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/domain/repository"
	"flyme-backend/app/interfaces/request"
	"flyme-backend/app/packages/auth"
)

type UserUseCase struct {
	dbRepository repository.DBRepositoryImpl
}

func NewUserUseCase(r repository.DBRepositoryImpl) *UserUseCase {
	return &UserUseCase{
		dbRepository: r,
	}
}

func (u *UserUseCase) ReadUser(userID string) (*entity.UserTable, error) {
	user, err := u.dbRepository.GetUser(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) CreateUser(req *request.CreateUserRequest) (*entity.InsertUser, error) {

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

	return query, nil
}

func (u *UserUseCase) UpdateUser(userID string, req *request.UpdateUserRequest) (*entity.PutUser, error) {

	query := &entity.PutUser{
		UserID:   userID,
		UserName: req.UserName,
		Icon:     req.Icon,
	}

	err := u.dbRepository.PutUser(query)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (u *UserUseCase) Login(req *request.LoginRequest) (string, error) {
	user, err := u.dbRepository.GetUser(req.UserID)

	if err != nil {
		return "", err
	}

	if req.Passwd != user.Passwd {
		return "", errors.New("password incorrect")
	}

	token, err := auth.GenerateUserToken(req.UserID, req.Passwd)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUseCase) ValidateUserToken(req *request.ValidateUserTokenRequest) error {

	_, err := auth.ValidateUserToken(req.Token)
	if err != nil {
		return err
	}

	return nil
}
