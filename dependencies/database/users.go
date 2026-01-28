package database

import (
	"core/app/domain"
	"core/app/domain/services/user"

	"gorm.io/gorm"
)

// UserRepositoryImpl implements the UserRepository interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (ur *UserRepositoryImpl) InsertNewUser(newUser domain.NewUser) error {
	// Implement the logic to insert a new user into the database

	user := MapToUserDB(newUser)

	result := ur.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
