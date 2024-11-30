package controllers

import (
	"fmt"
	"src/config"
	"src/models"
)

func AddBus(bus models.Bus) error {
	if result := config.DB.Create(&bus); result.Error != nil {
		return result.Error
	}
	fmt.Println("Bus added successfully")
	return nil
}

func GetAllBuses() ([]models.Bus, error) {
	var buses []models.Bus
	if result := config.DB.Find(&buses); result.Error != nil {
		return nil, result.Error
	}
	return buses, nil
}

// Add more functions for updating and deleting buses
