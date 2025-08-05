package app

import (
	"context"
	"core/app/domain/services"
	"core/dependencies/httpserver"
	"github.com/gin-gonic/gin"
)

type AppDependencies struct {
	HttpServer *gin.Engine
	Logger     services.Logger
	Ctx        context.Context
}

func (a AppDependencies) Start() error {

	// Create http server API
	router, addr, err := httpserver.CreateHttpServer(a.Logger)
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
