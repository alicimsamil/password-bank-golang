package model

import "time"

type Password struct {
	Id          int32
	Password    string
	Type        string
	Account     string
	ServiceName string
	Notes       string
	Date        time.Time
}
