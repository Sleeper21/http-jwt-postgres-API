package usecase

import (
	"core/app/domain"
	"core/app/domain/services"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context, logger services.Logger) (domain.NewUser, error) {
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

	return request, nil
}
