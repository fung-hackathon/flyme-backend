package entity

import "encoding/json"

type Deserialize interface {
	GetUser | GetTimeline | GetHistory | GetFollowers
}

type Serialize interface {
	InsertUser | HistoryTable
}

func BindToJsonStruct[T Deserialize](jm map[string]interface{}, js *T) error {
	jsonStr, err := json.Marshal(jm)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonStr, js)
	if err != nil {
		return err
	}

	return nil
}

func BindToJsonMap[T Serialize](js *T) (map[string]interface{}, error) {
	jsonStr, err := json.Marshal(js)
	if err != nil {
		return nil, err
	}

	var jm map[string]interface{}

	err = json.Unmarshal(jsonStr, &jm)
	if err != nil {
		return nil, err
	}

	return jm, nil
}
