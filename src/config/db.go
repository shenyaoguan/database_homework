package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"src/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:19171107@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Flight{}, &models.Hotel{}, &models.Bus{}, &models.Customer{}, &models.Reservation{})

	DB = database
}
