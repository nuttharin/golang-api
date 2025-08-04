package common

import (
	"time"
)

type Timestamp struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"index:,sort:desc" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
