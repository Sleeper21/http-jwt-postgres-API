package httpserver

import (
	"core/app/domain/services"
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
	router.Use(CheckAuthentication())
	router.GET("/health", deps.healthCheckHandler)
}

func (deps HttpRouter) healthCheckHandler(c *gin.Context) {
	deps.logger.Info("Health check: Status OK")

	c.JSON(200, gin.H{
		"server status": "Server is ok and running",
	})
}
