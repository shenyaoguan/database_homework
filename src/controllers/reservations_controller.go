package controllers

import (
	"fmt"
	"src/config"
	"src/models"
)

func AddReservation(reservation models.Reservation) error {
	if result := config.DB.Create(&reservation); result.Error != nil {
		return result.Error
	}
	fmt.Println("Reservation added successfully")
	return nil
}

func GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	if result := config.DB.Find(&reservations); result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

// Add more functions for updating and deleting reservations
