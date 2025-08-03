package database

import (
	model "golang-api/repositories/models"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	var count int64
	// Check count in table
	db.Model(&model.User{}).Count(&count)
	if count == 0 {
		users := []model.User{
			{Name: "user1", Email: "user1@test.com"},
			{Name: "user2", Email: "user2@test.com"},
		}

		return db.Create(&users).Error
	}
	return nil
}
