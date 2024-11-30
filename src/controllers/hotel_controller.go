package controllers

import (
	"fmt"
	"src/config"
	"src/models"
)

func AddHotel(hotel models.Hotel) error {
	if result := config.DB.Create(&hotel); result.Error != nil {
		return result.Error
	}
	fmt.Println("Hotel added successfully")
	return nil
}

func GetAllHotels() ([]models.Hotel, error) {
	var hotels []models.Hotel
	if result := config.DB.Find(&hotels); result.Error != nil {
		return nil, result.Error
	}
	return hotels, nil
}

// Add more functions for updating and deleting hotels
func UpdateHotel(name string, updatedHotel models.Hotel) error {
	var hotel models.Hotel
	if result := config.DB.First(&hotel, "name = ?", name); result.Error != nil {
		return result.Error
	}
	if result := config.DB.Model(&hotel).Updates(updatedHotel); result.Error != nil {
		return result.Error
	}
	fmt.Println("Hotel updated successfully")
	return nil
}

func DeleteHotel(name string) error {
	if result := config.DB.Delete(&models.Hotel{}, "name = ?", name); result.Error != nil {
		return result.Error
	}
	fmt.Println("Hotel deleted successfully")
	return nil
}
