package database

import (
	model "golang-api/repositories/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err

	}

	log.Println("Migration completed.")
	return nil
}
