package userService

import (
	"core/app/domain"
)

type UserRepository interface {
	InsertNewUser(newUser domain.NewUser) error
	GetUserByEmail(email string) (domain.NewUser, error)
	//GetAllUsers() ([]User, error)
	//UpdateUser(user User) error
	//DeleteUser(email string) error
}
