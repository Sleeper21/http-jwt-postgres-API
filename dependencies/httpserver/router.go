package httpserver

import (
	"core/app/domain/services"
	"core/app/usecase"

	"github.com/gin-gonic/gin"
)

type HttpRouter struct {
	logger services.Logger
	//authProvider
	//db
}

func CreateHttpRouter(logger services.Logger) HttpRouter {
	return HttpRouter{
		logger: logger,
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

		user, err := usecase.RegisterUser(ctx, deps.logger)
		if err != nil {
			return
		}

		ctx.JSON(200, gin.H{
			"message": "User registered successfully",
			"email":   user.Email,
		})
	}
}
