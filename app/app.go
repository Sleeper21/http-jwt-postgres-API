package app

import (
	"context"
	"core/app/domain/services"
	"core/app/domain/services/userService"
	"core/dependencies/httpserver"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppDependencies struct {
	HttpServer     *gin.Engine
	Logger         services.Logger
	Ctx            context.Context
	DB             *gorm.DB
	UserRepository userService.UserRepository // Repository for user operations
}

func (a AppDependencies) Start() error {

	// Create http server API
	router, addr, err := httpserver.CreateHttpServer(a.Logger, a.UserRepository)
	if err != nil {
		a.Logger.WithError(err, "Error creating HTTP server")
		panic(err)
	}

	// Store the router in the AppDependencies
	a.HttpServer = router

	// Run server
	err = httpserver.Run(router, addr)
	if err != nil {
		a.Logger.WithError(err, "Error starting HTTP server")
		return err
	}

	return nil
}
