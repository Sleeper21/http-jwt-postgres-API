package app

import (
	"context"
	"core/app/domain"
	"core/app/httpserver"
)

type AppDependencies struct {
	Logger domain.Logger
	Ctx    context.Context
}

func (a AppDependencies) Start() error {

	// Create http server API
	httpserver.CreateHttpServer(a.Logger)

	// Run server
	err := httpserver.Run()
	if err != nil {
		a.Logger.WithError(err, "Error starting HTTP server")
		return err
	}

	return nil
}
