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

func GetReservationByNum(reservationNum string) (models.Reservation, error) {
	var reservation models.Reservation
	if result := config.DB.First(&reservation, "reservation_num = ?", reservationNum); result.Error != nil {
		return models.Reservation{}, result.Error
	}
	return reservation, nil
}

// Add more functions for updating and deleting reservations
func UpdateReservation(reservationNum string, updatedReservation models.Reservation) error {
	var reservation models.Reservation
	if result := config.DB.First(&reservation, "reservation_num = ?", reservationNum); result.Error != nil {
		return result.Error
	}
	if result := config.DB.Model(&reservation).Updates(updatedReservation); result.Error != nil {
		return result.Error
	}
	fmt.Println("Reservation updated successfully")
	return nil
}

func DeleteReservation(reservationNum string) error {
	if result := config.DB.Delete(&models.Reservation{}, "reservation_num = ?", reservationNum); result.Error != nil {
		return result.Error
	}
	fmt.Println("Reservation deleted successfully")
	return nil

}
