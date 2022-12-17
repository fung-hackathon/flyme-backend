package usecase

import (
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/domain/repository"
	"flyme-backend/app/interfaces/request"
	"flyme-backend/app/interfaces/response"
)

type HistoryUseCase struct {
	dbRepository repository.DBRepositoryImpl
}

func NewHistoryUseCase(r repository.DBRepositoryImpl) *HistoryUseCase {
	return &HistoryUseCase{
		dbRepository: r,
	}
}

func (u *HistoryUseCase) StartHistory(userID string, req *request.StartHistoryRequest) (*response.StartHistoryResponse, error) {

	history, err := u.dbRepository.StartHistory(&entity.StartHistory{
		UserID:    userID,
		StartTime: req.StartTime,
	})

	if err != nil {
		return nil, err
	}

	rcoords := make([]response.Coordinate, len(history.Coords))

	for i, c := range history.Coords {
		rcoords[i] = response.Coordinate{
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
		}
	}

	res := &response.StartHistoryResponse{
		Coords: rcoords,
		Dist:   history.Dist,
		Finish: history.Finish,
		Start:  history.Start,
		State:  history.State,
	}

	return res, nil
}

func (u *HistoryUseCase) FinishHistory(userID string, req *request.FinishHistoryRequest) (*response.FinishHistoryResponse, error) {

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
		FinishTime: req.FinishTime,
	})

	if err != nil {
		return nil, err
	}

	rcoords := make([]response.Coordinate, len(req.Coords))

	for i, c := range history.Coords {
		rcoords[i] = response.Coordinate{
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
		}
	}

	res := &response.FinishHistoryResponse{
		Coords: rcoords,
		Dist:   history.Dist,
		Finish: history.Finish,
		Start:  history.Start,
		State:  history.State,
	}

	return res, nil
}

func (u *HistoryUseCase) ReadHistories(userID string, size int) (*response.ReadHistoriesResponse, error) {

	histories, err := u.dbRepository.GetHistories(userID, size)
	if err != nil {
		return nil, err
	}

	rtimeline := make([]response.HistoryTable, size)

	for i, history := range histories.Histories {

		rcoords := make([]response.Coordinate, len(history.Coords))

		for i, c := range history.Coords {
			rcoords[i] = response.Coordinate{
				Longitude: c.Longitude,
				Latitude:  c.Latitude,
			}
		}

		rtimeline[i] = response.HistoryTable{
			Coords: rcoords,
			Dist:   history.Dist,
			Finish: history.Finish,
			Start:  history.Start,
			State:  history.Start,
		}
	}

	res := &response.ReadHistoriesResponse{
		Histories: rtimeline,
	}

	return res, nil
}

func (u *HistoryUseCase) ReadTimeline(userID string, size int) (*response.ReadTimelineResponse, error) {
	timeline, err := u.dbRepository.GetTimeline(userID, size)
	if err != nil {
		return nil, err
	}

	rtimeline := make([]response.HistoryTimeline, size)

	for i, hid := range timeline.Histories {

		history, err := u.dbRepository.GetHistory(hid)
		if err != nil {
			return nil, err
		}

		user, err := u.dbRepository.GetUser(history.UserID)
		if err != nil {
			return nil, err
		}

		rtimeline[i] = response.HistoryTimeline{
			User: response.UserInfo{
				UserID:   user.UserID,
				UserName: user.UserName,
				Icon:     user.Icon,
			},
			Finish: history.Finish,
			Start:  history.Start,
			State:  history.State,
		}
	}

	res := &response.ReadTimelineResponse{
		Histories: rtimeline,
	}

	return res, nil
}
