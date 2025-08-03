package main

import (
	config "golang-api/configs"
	"golang-api/database"
	database_mysql "golang-api/database/mysql"
	"golang-api/httpserver"
	"golang-api/pkg/utils"
	"golang-api/routes"
	"log"
)

func main() {

	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal("cannot load config.json:", err)
	}

	isDebug := utils.StringToBool(config.Server.Debug)

	// ===== Set database =====

	var db database.Database
	mysqlDB := &database_mysql.MySQLDatabase{}
	if err := mysqlDB.Connect(config.Database); err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db = mysqlDB
	defer db.Close()

	// ===== End Set database =====

	// ===== Run Migration =====
	if err := database.Migrate(db.GetDB()); err != nil {
		panic("Migration failed: " + err.Error())
	}

	// ===== Run Seed Data =====
	if err := database.Seed(db.GetDB()); err != nil {
		panic("Seeding failed: " + err.Error())
	}

	router := httpserver.NewRouter(isDebug)
	routes.Setup(router, db.GetDB(), *config)

	host := config.Server.Host
	port := utils.StringToInt(config.Server.Port)

	// Set Gin
	restApi := httpserver.NewRestAPI(host, port, router)

	// Start Api
	if err := restApi.Start(); err != nil {
		panic("Failed to start server: " + err.Error())
	}

}
