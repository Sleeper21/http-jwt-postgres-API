package main

import (
	"context"
	"core/app"
	"core/app/domain/services"
	"core/dependencies/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log := logger.NewTextLogger()
	log.Infof("Starting execution on: %s\n", time.Now().String())

	ctx := context.Background()
	application := loadDependencies(log, ctx)

	err := application.Start()
	if err != nil {
		log.WithError(err, "Error starting application")
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down application...")

}

func loadDependencies(log services.Logger, ctx context.Context) app.AppDependencies {
	// Load dependencies here
	// For example, database connection, cache, etc.

	// Return the application dependencies
	return app.AppDependencies{
		Logger: log,
		Ctx:    ctx,
	}
}
