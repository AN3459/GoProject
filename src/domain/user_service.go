package domain

type UserService interface {
    AddUser(name string, age int) error
    GetUserByID(id int64) (*User, error)
}