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
	addr, err := httpserver.CreateHttpServer(a.Logger)
	if err != nil {
		a.Logger.WithError(err, "Error creating HTTP server")
		panic(err)
	}

	// Run server
	err = httpserver.Run(addr)
	if err != nil {
		a.Logger.WithError(err, "Error starting HTTP server")
		return err
	}

	return nil
}
