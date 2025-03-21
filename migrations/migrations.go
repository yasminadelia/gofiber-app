package migrations

import (
	"gofiber-app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID: "20250321_add_age_to_users",
            Migrate: func(tx *gorm.DB) error {
                return tx.AutoMigrate(&models.User{})
            },
            Rollback: func(tx *gorm.DB) error {
                return tx.Migrator().DropColumn(&models.User{}, "Age")
            },
        },
    })

    return m.Migrate()
}
