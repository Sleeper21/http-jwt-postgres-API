package httpserver

import (
	"core/app/domain/services"
	"core/app/domain/services/userService"
	"core/app/usecase/user"

	"github.com/gin-gonic/gin"
)

type HttpRouter struct {
	logger         services.Logger
	userRepository userService.UserRepository // Repository for user operations
	//authProvider
}

func CreateHttpRouter(logger services.Logger, userRepo userService.UserRepository) HttpRouter {
	return HttpRouter{
		logger:         logger,
		userRepository: userRepo,
	}

}

func (deps HttpRouter) SetRoutes(router *gin.Engine) {
	// status route
	router.GET("/health", healthCheckHandler(deps))
	router.POST("/register", registerNewUser(deps))

	// Authenticated routes
	authorized := router.Group("/")
	authorized.Use(CheckAuthentication())
}

func healthCheckHandler(deps HttpRouter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		deps.logger.Info("Health check: Status OK")

		ctx.JSON(200, gin.H{
			"server status": "Server is ok and running",
		})
	}
}

func registerNewUser(deps HttpRouter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Call the use case with repository injected
		user, err := user.RegisterUser(ctx, deps.logger, deps.userRepository)
		if err != nil {
			deps.logger.WithError(err, "Error registering new user")
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "User registered successfully",
			"email":   user.Email,
		})
	}
}
