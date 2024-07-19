package service

import (
	"password-bank-golang/data/repository"
	"password-bank-golang/service/model"
)

type IPasswordService interface {
	GetPasswordById(id string, email string) (model.Password, error)
	GetAllPasswords(email string) ([]model.Password, error)
	CreatePassword(password model.Password, email string) error
	UpdatePassword(password model.Password, email string) error
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

func (service *PasswordService) CreatePassword(password model.Password, email string) error {
	return service.repository.CreatePassword(password, email)
}

func (service *PasswordService) UpdatePassword(password model.Password, email string) error {
	return service.repository.UpdatePassword(password, email)
}

func (service *PasswordService) DeletePassword(id string, email string) error {
	return service.repository.DeletePassword(id, email)
}

func NewPasswordService(repo repository.IPasswordRepository) IPasswordService {
	return &PasswordService{repository: repo}
}
