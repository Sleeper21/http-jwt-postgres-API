package main

import (
	"context"
	"core/app"
	"core/dependencies/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	log := logger.NewTextLogger()
	log.Infof("Starting execution on: %s\n", time.Now().String())

	ctx := context.Background()
	application := loadDependencies(log, ctx)

	err = application.Start()
	if err != nil {
		log.WithError(err, "Error starting application")
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down application...")

}

func loadDependencies(logger logger.LmptLogger, ctx context.Context) app.AppDependencies {
	// Load dependencies here
	// For example, database connection, cache, etc.

	// Return the application dependencies
	return app.AppDependencies{
		Logger: logger,
		Ctx:    ctx,
	}
}
