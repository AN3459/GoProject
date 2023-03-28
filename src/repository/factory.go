package repository

import (
	"example.com/m/v2/src/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user *entity.User) error {
	return r.db.Create(user).Error
}
