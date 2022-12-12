package entity

import "encoding/json"

type GetUser struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}

type InsertUser struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Passwd   string `json:"passwd"`
	Icon     string `json:"icon"`
}

type PutUser struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
	Icon     string `json:"icon"`
}

type Deserialize interface {
	GetUser
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

type Serialize interface {
	InsertUser
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
