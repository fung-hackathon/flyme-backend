package usecase

import (
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/domain/repository"
	"flyme-backend/app/interfaces/request"
)

type HistoryUseCase struct {
	dbRepository repository.DBRepositoryImpl
}

func NewHistoryUseCase(r repository.DBRepositoryImpl) *HistoryUseCase {
	return &HistoryUseCase{
		dbRepository: r,
	}
}

func (u *HistoryUseCase) StartHistory(userID string, req *request.StartHistoryRequest) (*entity.HistoryTable, error) {

	history, err := u.dbRepository.StartHistory(&entity.StartHistory{
		UserID:    userID,
		Ticket:    req.Ticket,
		StartTime: req.StartTime,
	})

	if err != nil {
		return nil, err
	}

	return history, nil
}

func (u *HistoryUseCase) FinishHistory(userID string, req *request.FinishHistoryRequest) (*entity.HistoryTable, error) {

	hcoords := make([]entity.Coordinate, len(req.Coords))

	for i, c := range req.Coords {
		hcoords[i] = entity.Coordinate{
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
		}
	}

	history, err := u.dbRepository.FinishHistory(&entity.FinishHistory{
		Coords:     hcoords,
		UserID:     userID,
		Distance:   req.Distance,
		FinishTime: req.FinishTime,
	})

	if err != nil {
		return nil, err
	}

	return history, nil
}

func (u *HistoryUseCase) ReadHistories(userID string, size int) (*entity.GetHistories, error) {

	histories, err := u.dbRepository.GetHistories(userID, size)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (u *HistoryUseCase) ReadTimeline(userID string, size int) ([]*entity.GetHistory, []*entity.GetUser, error) {
	timeline, err := u.dbRepository.GetTimeline(userID, size)
	if err != nil {
		return nil, nil, err
	}

	histories := []*entity.GetHistory{}
	users := []*entity.GetUser{}

	for i, hid := range timeline.Histories {

		if i >= int(size) {
			break
		}

		history, err := u.dbRepository.GetHistory(hid)
		if err != nil {
			return nil, nil, err
		}

		histories = append(histories, history)

		user, err := u.dbRepository.GetUser(history.UserID)
		if err != nil {
			return nil, nil, err
		}

		users = append(users, user)
	}

	return histories, users, nil
}
