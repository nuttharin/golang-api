package database_mysql

import (
	"fmt"
	"golang-api/database"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gorm.io/plugin/dbresolver"
)

type MySQLDatabase struct {
	db *gorm.DB
}

func (m *MySQLDatabase) Connect(config database.DatabaseConnection) error {

	log.Println(m.buildDBConnection(config.Write))

	db, err := gorm.Open(mysql.Open(m.buildDBConnection(config.Write)))
	if err != nil {
		return err
	}

	err = db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.Open(m.buildDBConnection(config.Read))},
			Sources:  []gorm.Dialector{mysql.Open(m.buildDBConnection(config.Write))},
			Policy:   dbresolver.RandomPolicy{},
		}).

		// Set Config
		SetConnMaxLifetime(5 * time.Minute).
		SetMaxIdleConns(10).
		SetMaxOpenConns(100))

	if err != nil {
		return err
	}

	// Set connected database instance
	m.db = db

	return nil
}

func (m *MySQLDatabase) Close() error {
	if m.db != nil {
		sqlDB, _ := m.db.DB()
		sqlDB.Close()
	}

	return nil
}

func (m *MySQLDatabase) GetDB() *gorm.DB {
	return m.db
}

func (m *MySQLDatabase) buildDBConnection(db database.DatabaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		db.User, db.Pass, db.Host, db.Port, db.Name,
	)
}
