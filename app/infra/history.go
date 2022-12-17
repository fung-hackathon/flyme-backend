package infra

import (
	"errors"
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/packages/geo"
	"sort"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func (r *DBRepository) GetHistory(historyID string) (*entity.GetHistory, error) {
	doc := r.Client.Collection("histories").Doc(historyID)

	exist, err := r.checkIfDataExists(doc)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, ErrHistoryNotFound
	}

	docSnap, err := doc.Get(r.Context)
	if err != nil {
		return nil, err
	}

	var history entity.GetHistory
	err = entity.BindToJsonStruct(docSnap.Data(), &history)
	if err != nil {
		return nil, err
	}

	return &history, nil
}

func (r *DBRepository) GetHistories(userID string, size int) (*entity.GetHistories, error) {
	iter := r.Client.Collection("histories").Where("userID", "==", userID).Documents(r.Context)

	histories := []entity.GetHistory{}

	for {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var history entity.GetHistory
		err = entity.BindToJsonStruct(docSnap.Data(), &history)
		if err != nil {
			return nil, err
		}

		histories = append(histories, history)

		if err != nil {
			return nil, err
		}
	}

	// Future Work: sortとlimitはデータ取得時に行うべき
	sort.Slice(histories, func(i, j int) bool {
		return histories[i].Start > histories[j].Start
	})

	if len(histories) >= size {
		histories = histories[:size]
	}

	return &entity.GetHistories{
		Histories: histories,
	}, nil
}

func (r *DBRepository) StartHistory(history *entity.StartHistory) (*entity.HistoryTable, error) {

	historyID := entity.NewHistoryID()
	historyTable := entity.HistoryTable{
		Coords:    make([]entity.Coordinate, 0),
		Dist:      0.,
		Finish:    "",
		Start:     history.StartTime,
		State:     "start",
		Ticket:    history.Ticket,
		UserID:    history.UserID,
		HistoryID: historyID,
	}

	// userのhistoryIDInProgressの更新
	{
		doc := r.Client.Collection("users").Doc(historyTable.UserID)

		exist, err := r.checkIfDataExists(doc)
		if err != nil {
			return nil, err
		}
		if !exist {
			return nil, ErrUserNotFound
		}

		docSnap, err := doc.Get(r.Context)
		if err != nil {
			return nil, err
		}

		var user entity.GetUser
		err = entity.BindToJsonStruct(docSnap.Data(), &user)
		if err != nil {
			return nil, err
		}

		if user.HistoryIDInProgress != "" {
			_, err := r.Client.Collection("histories").Doc(user.HistoryIDInProgress).Delete(r.Context)
			if err != nil {
				return nil, err
			}
		}

		info := []firestore.Update{
			{Path: "historyIDInProgress", Value: historyTable.HistoryID},
		}

		_, err = doc.Update(r.Context, info)
		if err != nil {
			return nil, err
		}
	}

	// historiesにhistoryを追加
	{
		doc := r.Client.Collection("histories").Doc(historyTable.HistoryID)

		data, err := entity.BindToJsonMap(&historyTable)
		if err != nil {
			return nil, err
		}

		_, err = doc.Set(r.Context, data)
		if err != nil {
			return nil, err
		}
	}

	return &historyTable, nil
}

func (r *DBRepository) FinishHistory(history *entity.FinishHistory) (*entity.HistoryTable, error) {

	// userのドキュメントを保持
	userDoc := r.Client.Collection("users").Doc(history.UserID)

	exist, err := r.checkIfDataExists(userDoc)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, ErrUserNotFound
	}

	// user情報の取得
	var user entity.GetUser
	{
		docSnap, err := userDoc.Get(r.Context)
		if err != nil {
			return nil, err
		}

		err = entity.BindToJsonStruct(docSnap.Data(), &user)
		if err != nil {
			return nil, err
		}
	}

	historyID := user.HistoryIDInProgress

	if historyID == "" {
		return nil, errors.New("unknown history")
	}

	geoCoords := make([]geo.Coordinate, len(history.Coords))
	for i, c := range history.Coords {
		geoCoords[i] = geo.Coordinate{
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
		}
	}

	// 総距離を計測
	// dist, err := geo.GetDistanceKm(geoCoords)
	// if err != nil {
	// 	return nil, err
	// }

	// historyのドキュメントを保持
	historiesDoc := r.Client.Collection("histories").Doc(historyID)

	// 更新情報を一旦取りまとめる
	historyTable := entity.HistoryTable{
		Coords:    history.Coords,
		Dist:      history.Distance,
		Finish:    history.FinishTime,
		State:     "finish",
		UserID:    history.UserID,
		HistoryID: historyID,
	}

	exist, err = r.checkIfDataExists(historiesDoc)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, ErrHistoryNotFound
	}

	// historyを更新
	{
		info := []firestore.Update{
			{Path: "state", Value: historyTable.State},
			{Path: "coordinates", Value: historyTable.Coords},
			{Path: "dist", Value: historyTable.Dist},
			{Path: "finish", Value: historyTable.Finish},
		}

		_, err = historiesDoc.Update(r.Context, info)
		if err != nil {
			return nil, err
		}
	}

	// userのhistoryIDInProgressの更新
	{
		info := []firestore.Update{
			{Path: "historyIDInProgress", Value: ""},
		}

		_, err = userDoc.Update(r.Context, info)
		if err != nil {
			return nil, err
		}
	}

	// 全followersのTLを更新
	{
		followers, err := r.GetFollowers(historyTable.UserID)
		if err != nil {
			return nil, err
		}

		for _, fid := range followers.Followers {
			_, err := r.insertHistoryToTimeline(fid, historyTable.HistoryID)
			if err != nil {
				return nil, err
			}
		}
	}

	// 以下, 更新したhistoryをGetして返す
	docSnap, err := historiesDoc.Get(r.Context)
	if err != nil {
		return nil, err
	}

	var historyResult entity.HistoryTable
	err = entity.BindToJsonStruct(docSnap.Data(), &historyResult)
	if err != nil {
		return nil, err
	}

	// 頼む！動いてくれ！
	return &historyResult, nil
}
