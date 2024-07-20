package model

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PasswordRequest struct {
	Id          int32  `json:"id"`
	Password    string `json:"password"`
	Type        string `json:"type"`
	Account     string `json:"account"`
	ServiceName string `json:"service_name"`
	Notes       string `json:"notes"`
}
