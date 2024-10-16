package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"backend/models"
	"log"
	"backend/config"
)

var DB *gorm.DB

func Connect() {
	dsn := config.GetEnv("DB_DSN")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
}