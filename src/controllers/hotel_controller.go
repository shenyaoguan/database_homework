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

func GetHotelByLocation(location string) (models.Hotel, error) {
	var hotel models.Hotel
	if result := config.DB.First(&hotel, "location = ?", location); result.Error != nil {
		return models.Hotel{}, result.Error
	}
	return hotel, nil
}

func GetHotelsByPrice(price int) ([]models.Hotel, error) {
	var hotels []models.Hotel
	if result := config.DB.Find(&hotels, "price = ?", price); result.Error != nil {
		return nil, result.Error
	}
	return hotels, nil
}

func GetHotelsByNumRooms(numRooms int) ([]models.Hotel, error) {
	var hotels []models.Hotel
	if result := config.DB.Find(&hotels, "num_rooms = ?", numRooms); result.Error != nil {
		return nil, result.Error
	}
	return hotels, nil
}

func GetHotelsByNumAvail(numAvail int) ([]models.Hotel, error) {
	var hotels []models.Hotel
	if result := config.DB.Find(&hotels, "num_avail = ?", numAvail); result.Error != nil {
		return nil, result.Error
	}
	return hotels, nil
}

func GetHotelsByLocationAndPrice(location string, price int) ([]models.Hotel, error) {
	var hotels []models.Hotel
	if result := config.DB.Find(&hotels, "location = ? AND price = ?", location, price); result.Error != nil {
		return nil, result.Error
	}
	return hotels, nil
}

// Add more functions for updating and deleting hotels
func UpdateHotel(name string, updatedHotel models.Hotel) error {
	var hotel models.Hotel
	if result := config.DB.First(&hotel, "location = ?", name); result.Error != nil {
		return result.Error
	}
	if result := config.DB.Model(&hotel).Updates(updatedHotel); result.Error != nil {
		return result.Error
	}
	fmt.Println("Hotel updated successfully")
	return nil
}

func DeleteHotel(name string) error {
	if result := config.DB.Delete(&models.Hotel{}, "location = ?", name); result.Error != nil {
		return result.Error
	}
	fmt.Println("Hotel deleted successfully")
	return nil
}
