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

func BindToJson[T Deserialize](data map[string]interface{}, js *T) error {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonStr, js)
	if err != nil {
		return err
	}

	return nil
}
