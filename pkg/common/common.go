package common

import (
	"time"

	"gorm.io/gorm"
)

type Timestamp struct {
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `gorm:"index:,sort:desc" json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
