package model

type Students struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	RegisteredDate string `json:"registered-date"`
	Address        string `json:"address"`
}
