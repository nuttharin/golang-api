package database

import (
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	Write DatabaseConfig
	Read  DatabaseConfig
}

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type Database interface {
	Connect(config DatabaseConnection) error
	Close() error
	GetDB() *gorm.DB
}
