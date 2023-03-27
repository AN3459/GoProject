package domain

type UserRepository interface {
    AddUser(user *User) error
    GetUserByID(id int64) (*User, error)
}