package migrations

import (
	"gofiber-app/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}) // Creates or updates the `users` table
}

