package database

import "core/app/domain"

func MapToUserDB(user domain.NewUser) User {
	return User{
		Email:    user.Email,
		Password: user.Password,
	}
}
