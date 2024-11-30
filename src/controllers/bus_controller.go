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

func GetBusByLocation(location string) (models.Bus, error) {
	var bus models.Bus
	if result := config.DB.First(&bus, "location = ?", location); result.Error != nil {
		return models.Bus{}, result.Error
	}
	return bus, nil
}

func UpdateBus(location string, updatedBus models.Bus) error {
	var bus models.Bus
	if result := config.DB.First(&bus, "location = ?", location); result.Error != nil {
		return result.Error
	}
	if result := config.DB.Model(&bus).Updates(updatedBus); result.Error != nil {
		return result.Error
	}
	fmt.Println("Bus updated successfully")
	return nil
}

func DeleteBus(location string) error {
	if result := config.DB.Delete(&models.Bus{}, "location = ?", location); result.Error != nil {
		return result.Error
	}
	fmt.Println("Bus deleted successfully")
	return nil
}
