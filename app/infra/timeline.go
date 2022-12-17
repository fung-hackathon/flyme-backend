package infra

import (
	"flyme-backend/app/domain/entity"

	"cloud.google.com/go/firestore"
)

func (r *DBRepository) GetTimeline(userID string, size int) (*entity.GetTimeline, error) {
	doc := r.Client.Collection("timelines").Doc(userID)

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

	var timeline entity.GetTimeline
	err = entity.BindToJsonStruct(docSnap.Data(), &timeline)
	if err != nil {
		return nil, err
	}

	if len(timeline.Histories) > size {
		timeline.Histories = timeline.Histories[:size]
	}

	return &timeline, nil
}

func (r *DBRepository) insertHistoryToTimeline(userID string, historyID string) (*entity.TimelineTable, error) {
	doc := r.Client.Collection("timelines").Doc(userID)

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

	var timeline entity.GetTimeline
	err = entity.BindToJsonStruct(docSnap.Data(), &timeline)
	if err != nil {
		return nil, err
	}

	timeline.Histories, timeline.Histories[0] = append(timeline.Histories[:1], timeline.Histories[0:]...), historyID

	info := []firestore.Update{
		{Path: "histories", Value: timeline.Histories},
	}

	_, err = doc.Update(r.Context, info)
	if err != nil {
		return nil, err
	}

	return &timeline, nil
}
