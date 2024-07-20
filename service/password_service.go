package service

import (
	passModel "password-bank-golang/api/controller/model"
	"password-bank-golang/data/repository"
	"password-bank-golang/service/model"
	"time"
)

type IPasswordService interface {
	GetPasswordById(id string, email string) (model.Password, error)
	GetAllPasswords(email string) ([]model.Password, error)
	InsertPassword(password passModel.PasswordRequest, email string) error
	UpdatePassword(password passModel.PasswordRequest, email string) error
	DeletePassword(id string, email string) error
}

type PasswordService struct {
	repository repository.IPasswordRepository
}

func (service *PasswordService) GetPasswordById(id string, email string) (model.Password, error) {
	return service.repository.GetPasswordById(id, email)
}

func (service *PasswordService) GetAllPasswords(email string) ([]model.Password, error) {
	return service.repository.GetAllPasswords(email)
}

func (service *PasswordService) InsertPassword(password passModel.PasswordRequest, email string) error {
	return service.repository.InsertPassword(model.Password{
		Password:    password.Password,
		Type:        password.Type,
		Account:     password.Account,
		ServiceName: password.ServiceName,
		Notes:       password.Notes,
		Date:        time.Now(),
	}, email)
}

func (service *PasswordService) UpdatePassword(password passModel.PasswordRequest, email string) error {
	return service.repository.UpdatePassword(model.Password{
		Id:          password.Id,
		Password:    password.Password,
		Type:        password.Type,
		Account:     password.Account,
		ServiceName: password.ServiceName,
		Notes:       password.Notes,
		Date:        time.Now(),
	}, email)
}

func (service *PasswordService) DeletePassword(id string, email string) error {
	return service.repository.DeletePassword(id, email)
}

func NewPasswordService(repo repository.IPasswordRepository) IPasswordService {
	return &PasswordService{repository: repo}
}
