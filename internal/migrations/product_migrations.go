package migrations

import (
	"github.com/AIhmed/go-api-test/internal/models"
	"gorm.io/gorm"
)

func MigrateProducts(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.Inventory{},
		// Other product-related models...
	)
}
