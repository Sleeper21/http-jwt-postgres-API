package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Email     string         `gorm:"uniqueIndex;not null"`
	Password  string         `gorm:"not null"`
	CreatedAt int64          `gorm:"autoCreateTime"`
	UpdatedAt int64          `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate is a GORM hook that runs before inserting a record
// It automatically generates a UUID for new users
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
