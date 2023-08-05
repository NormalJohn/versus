package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect to PostgresSQL
func Connect() bool {
	var err error
	dsn := "host=localhost port=5432 user=postgres dbname=versus password=example sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Users{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
	return true
}
