package repository

import "example.com/m/v2/src/entity"

type UserRepository interface {
	Save(user *entity.User) error
}
