package main

import (
	"context"
	"core/app"
	"core/app/domain/services"
	"core/dependencies/database"
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
	// Connect to database
	db, err := database.Connect(log)
	if err != nil {
		log.WithError(err, "Failed to connect to database")
		panic(err)
	}

	// Run database migrations
	err = database.RunMigrations(db, log)
	if err != nil {
		log.WithError(err, "Failed to run migrations")
		panic(err)
	}

	// Create UserRepository with database connection
	// This is where we "inject" the dependency
	userRepo := database.NewUserRepository(db)
	log.Info("UserRepository initialized")

	// Return the application dependencies
	return app.AppDependencies{
		Logger:         log,
		Ctx:            ctx,
		DB:             db,
		UserRepository: userRepo,
	}
}
