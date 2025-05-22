package migrations

import (
	"gorm.io/gorm"
	"log"
)

type MigrationFunc func(*gorm.DB) error

var allMigrations = []MigrationFunc{
	MigrateUsers,
	MigrateProducts,
	// Add other migration functions here...
}

func RunAll(db *gorm.DB) error {
	log.Println("Starting database migrations...")
	for i, migration := range allMigrations {
		log.Printf("Running migration %d/%d", i+1, len(allMigrations))
		if err := migration(db); err != nil {
			log.Printf("Migration failed: %v", err)
			return err
		}
	}
	log.Println("Migrations completed successfully")
	return nil
}
