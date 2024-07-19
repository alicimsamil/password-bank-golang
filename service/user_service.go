package service

import "password-bank-golang/data/repository"

type IUserService interface {
	LoginUser(email string, password string) error
	CreateUser(email string, password string) error
}

type UserService struct {
	repo repository.IUserRepository
}

func (service *UserService) LoginUser(email string, password string) error {
	return service.repo.LoginUser(email, password)
}

func (service *UserService) CreateUser(email string, password string) error {
	return service.repo.CreateUser(email, password)
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}
