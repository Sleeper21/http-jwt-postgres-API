package user

import (
	"core/app/domain"
	"errors"
)

func ValidateInput(request domain.NewUser) (bool, error) {
	if request.Email == "" || request.Password == "" {
		return false, errors.New("email and password can't be empty")
	}

	return true, nil
}
