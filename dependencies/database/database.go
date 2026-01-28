package database

import (
	"core/app/domain/services"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect creates a connection to the PostgreSQL database using GORM
func Connect(logger services.Logger) (*gorm.DB, error) {
	// Build connection string from environment variables
	host := os.Getenv("DB_USERS_HOST")
	port := os.Getenv("DB_USERS_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	logger.Infof("Connecting to database: %s@%s:%s/%s", user, host, port, dbname)

	// Open connection with GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.WithError(err, "Failed to connect to database")
		return nil, err
	}

	logger.Info("Database connection established successfully")
	return db, nil
}
