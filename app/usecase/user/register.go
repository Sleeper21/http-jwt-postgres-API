package user

import (
	"core/app/domain"
	"core/app/domain/services"
	"core/app/domain/services/userService"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

// RegisterUser handles the business logic for user registration
func RegisterUser(ctx *gin.Context, logger services.Logger, userRepo userService.UserRepository) (domain.NewUser, error) {
	var request domain.NewUser

	err := ctx.ShouldBind(&request)
	if err != nil {
		logger.WithError(err, "Error binding JSON while registering new user")
		if strings.Contains(err.Error(), "EOF") {
			return domain.NewUser{}, errors.New("request body is empty")
		}
		if strings.Contains(err.Error(), "binding") {
			return domain.NewUser{}, errors.New("invalid request data")
		}
		if strings.Contains(err.Error(), "email") {
			return domain.NewUser{}, errors.New("invalid email")
		}
		return domain.NewUser{}, err
	}

	// Sanitize email input
	request.Email = strings.TrimSpace(request.Email)

	// Validate input request
	isValid, err := ValidateInput(request)
	if err != nil {
		logger.WithError(err, "Error validating request input")
		return domain.NewUser{}, err
	}

	if !isValid {
		logger.Error("Invalid input data for user registration")
		return domain.NewUser{}, errors.New("invalid input data")
	}

	// Check if user already exists
	found, err := userRepo.GetUserByEmail(request.Email)
	if err != nil {
		logger.WithError(err, "Error checking existing user in database")
		return domain.NewUser{}, err
	}

	if found.Email != "" {
		logger.Errorf("User with email %s already exists", request.Email)
		return domain.NewUser{}, errors.New("user email already exists")
	}

	logger.Infof("Registering new user: %s", request.Email)

	// Use repository to insert the user
	err = userRepo.InsertNewUser(request)
	if err != nil {
		logger.WithError(err, "Error inserting new user into database")
		return domain.NewUser{}, err
	}

	return request, nil
}
