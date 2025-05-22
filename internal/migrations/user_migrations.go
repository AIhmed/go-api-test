package migrations

import (
	"github.com/AIhmed/go-api-test/internal/models"
	"gorm.io/gorm"
)

func MigrateUsers(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Role{},
		// Other user-related models...
	)
}
