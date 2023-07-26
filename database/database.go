package database

import (
	"api/migrations"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=5435 user=postgres dbname=api_golang sslmode=disable password=admin"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error", err)
	}

	db = database
	migrations.RunMigrations(db)
	config, _ := db.DB()
	config.SetConnMaxIdleTime(10)
	config.SetMaxIdleConns(100)
	config.SetConnMaxIdleTime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}
