package usecase

import (
	"errors"
	"flyme-backend/app/domain/repository"
	"flyme-backend/app/logger"
	"image"

	"io"
)

type ImageUseCase struct {
	imageRepository repository.ImageRepositoryImpl
	userRepository  repository.DBRepositoryImpl
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrImageNotPNG  = errors.New("receive no png file")
)

func NewImageUseCase(ir repository.ImageRepositoryImpl, dr repository.DBRepositoryImpl) *ImageUseCase {
	return &ImageUseCase{ir, dr}
}

func (u *ImageUseCase) ValidateImg(file io.Reader) error {
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		logger.Log{
			Message: "validate image",
			Cause:   err,
		}.Err()
		return err
	}
	if format != "png" {
		logger.Log{
			Message: "validate image",
			Cause:   ErrImageNotPNG,
		}.Err()
		return ErrImageNotPNG
	}
	return nil
}

func (u *ImageUseCase) UploadIconImg(file io.Reader, userID string) error {
	ok, err := u.userRepository.CheckUserExist(userID)
	if err != nil {
		return err
	}
	if !ok {
		return ErrUserNotFound
	}
	return u.imageRepository.UploadIconImg(file, userID)
}

func (u *ImageUseCase) DownloadIconImg(file io.Writer, userID string) error {
	ok, err := u.userRepository.CheckUserExist(userID)
	if err != nil {
		return err
	} else if !ok {
		return ErrUserNotFound
	}

	err = u.imageRepository.DownloadIconImg(file, userID)
	if err != nil {
		return err
	}

	return nil
}
