package service

import (
	"example.com/m/v2/src/entity"
	"example.com/m/v2/src/repository"
)

type UserService interface {
	AddUser(name, email, password string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) AddUser(name, email, password string) error {
	user := &entity.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return s.repo.Save(user)
}
