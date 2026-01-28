package usecase

import (
	"core/app/domain"
	"core/app/domain/services"
	"core/app/domain/services/user"

	"github.com/gin-gonic/gin"
)

// RegisterUser handles the business logic for user registration
func RegisterUser(ctx *gin.Context, logger services.Logger, userRepo user.UserRepository) (domain.NewUser, error) {
	var request domain.NewUser

	err := ctx.ShouldBind(&request)
	if err != nil {
		logger.WithError(err, "Error binding JSON while registering new user")
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return domain.NewUser{}, err
	}

	if request.Email == "" || request.Password == "" {
		logger.Error("error. email and password cant be empty")
		ctx.JSON(400, gin.H{"error": "email and password cant be empty"})
		return domain.NewUser{}, nil
	}

	logger.Infof("Registering new user: %s", request.Email)
	logger.Infof("User Password: %s", request.Password)

	// Use repository to insert the user
	err = userRepo.InsertNewUser(request)
	if err != nil {
		logger.WithError(err, "Error inserting new user into database")
		ctx.JSON(500, gin.H{"error": "internal server error"})
		return domain.NewUser{}, err
	}

	return request, nil
}
