package database

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Question{},
		&models.Option{},
		&models.GameSession{},
	)
	if err != nil {
		return err
	}
	return nil
}
