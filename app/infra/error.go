package infra

import "errors"

var (
	ErrFirestoreConnection = errors.New("failed to establish connection to Firestore")
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrHistoryNotFound     = errors.New("history not found")
)
