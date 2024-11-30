package controllers

import (
	"fmt"
	"src/config"
	"src/models"
)

func AddFlight(flight models.Flight) error {
	if result := config.DB.Create(&flight); result.Error != nil {
		return result.Error
	}
	fmt.Println("Flight added successfully")
	return nil
}

func GetAllFlights() ([]models.Flight, error) {
	var flights []models.Flight
	if result := config.DB.Find(&flights); result.Error != nil {
		return nil, result.Error
	}
	return flights, nil
}

func GetFlightByNum(flightNum string) (models.Flight, error) {
	var flight models.Flight
	if result := config.DB.First(&flight, "flight_num = ?", flightNum); result.Error != nil {
		return models.Flight{}, result.Error
	}
	return flight, nil
}

func DeleteFlight(flightNum string) error {
	if result := config.DB.Delete(&models.Flight{}, "flight_num = ?", flightNum); result.Error != nil {
		return result.Error
	}
	fmt.Println("Flight deleted successfully")
	return nil

}

func UpdateFlight(flightNum string, updatedFlight models.Flight) error {
	var flight models.Flight
	if result := config.DB.First(&flight, "flight_num = ?", flightNum); result.Error != nil {
		return result.Error
	}
	if result := config.DB.Model(&flight).Updates(updatedFlight); result.Error != nil {
		return result.Error
	}
	fmt.Println("Flight updated successfully")
	return nil
}
