package model

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
