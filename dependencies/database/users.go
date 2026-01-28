package database

import (
	"core/app/domain"
	"core/app/domain/services/userService"

	"gorm.io/gorm"
)

// UserRepositoryImpl implements the UserRepository interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) userService.UserRepository {
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

func (ur *UserRepositoryImpl) GetUserByEmail(email string) (domain.NewUser, error) {
	// Find email in db and return it
	var user User
	result := ur.db.Where("email = ?", email).First(&user)
	if result.Error != nil {

		return domain.NewUser{}, result.Error
	}

	return domain.NewUser{
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
