package models

import (
	"golang-api/pkg/common"
)

type User struct {
	Id    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(64)" json:"name"`
	Email string `gorm:"type:varchar(64)" json:"email"`
	common.Timestamp
}
