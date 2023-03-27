package infrastructure

import (
    "example.com/m/v2/src/domain"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type mysqlUserRepository struct {
    db *gorm.DB
}

func NewUserRepo(db *gorm.DB) domain.UserRepository {
    return &mysqlUserRepository{db}
}

func (r *mysqlUserRepository) AddUser(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *mysqlUserRepository) GetUserByID(id int64) (*domain.User, error) {
    var user domain.User
    if err := r.db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
