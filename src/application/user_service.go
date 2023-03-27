package application

import "../domain"

type userService struct {
    userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
    return &userService{userRepo}
}

func (s *userService) AddUser(name string, age int) error {
    user := &domain.User{Name: name, Age: age}
    return s.userRepo.AddUser(user)
}

func (s *userService) GetUserByID(id int64) (*domain.User, error) {
    return s.userRepo.GetUserByID(id)
}
