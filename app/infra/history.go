package infra

import (
	"flyme-backend/app/domain/entity"
	"flyme-backend/app/packages/geo"

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
	iter := r.Client.Collection("histories").Where("user_id", "==", userID).Documents(r.Context)

	histories := make([]entity.GetHistory, size)

	for i := 0; i < size; i++ {
		docSnap, err := iter.Next()
		if err == iterator.Done {
			break
		}

		var history entity.GetHistory
		err = entity.BindToJsonStruct(docSnap.Data(), &history)
		if err != nil {
			return nil, err
		}

		histories[i] = history

		if err != nil {
			return nil, err
		}
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
		UserID:    history.UserID,
		HistoryID: historyID,
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

		info := []firestore.Update{
			{Path: "historyIDInProgress", Value: historyTable.HistoryID},
		}

		_, err = doc.Update(r.Context, info)
		if err != nil {
			return nil, err
		}
	}

	return &historyTable, nil
}

func (r *DBRepository) FinishHistory(history *entity.FinishHistory) (*entity.HistoryTable, error) {

	var user entity.GetUser

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

	geoCoords := make([]geo.Coordinate, len(history.Coords))
	for i, c := range history.Coords {
		geoCoords[i] = geo.Coordinate{
			Longitude: c.Longitude,
			Latitude:  c.Latitude,
		}
	}

	// 総距離を計測
	dist, err := geo.GetDistanceKm(geoCoords)
	if err != nil {
		return nil, err
	}

	// 更新情報を一旦取りまとめる
	historyTable := entity.HistoryTable{
		Coords:    history.Coords,
		Dist:      dist,
		Finish:    history.FinishTime,
		State:     "finish",
		UserID:    history.UserID,
		HistoryID: historyID,
	}

	// historyのドキュメントを保持
	historiesDoc := r.Client.Collection("histories").Doc(historyTable.HistoryID)

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
			{Path: "coordinates", Value: historyTable.Coords},
			{Path: "dist", Value: historyTable.Dist},
			{Path: "finish", Value: historyTable.Finish},
			{Path: "state", Value: historyTable.State},
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
		followers, err := r.GetFollowers(historyTable.HistoryID)
		if err != nil {
			return nil, err
		}

		for _, fid := range followers.Followers {
			r.insertHistoryToTimeline(fid, historyTable.HistoryID)
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
