package model

import "time"

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PasswordRequest struct {
	Id          int32     `json:"token"`
	Password    string    `json:"password"`
	Type        string    `json:"type"`
	Account     string    `json:"account"`
	ServiceName string    `json:"service_ame"`
	Notes       string    `json:"notes"`
	Date        time.Time `json:"date"`
}
