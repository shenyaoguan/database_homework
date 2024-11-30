package controllers

import (
	"fmt"
	"src/config"
	"src/models"
)

func AddReservation(reservation models.Reservation) error {
	var available int

	switch reservation.ResvType {
	case 1: // Flight
		var flight models.Flight
		if result := config.DB.First(&flight, "flight_num = ?", reservation.ResvKey); result.Error != nil {
			return result.Error
		}
		available = flight.NumAvail
		if available > 0 {
			flight.NumAvail--
			config.DB.Save(&flight)
		} else {
			return fmt.Errorf("no available seats on flight")
		}
	case 2: // Hotel
		var hotel models.Hotel
		if result := config.DB.First(&hotel, "location = ?", reservation.ResvKey); result.Error != nil {
			return result.Error
		}
		available = hotel.NumAvail
		if available > 0 {
			hotel.NumAvail--
			config.DB.Save(&hotel)
		} else {
			return fmt.Errorf("no available rooms in hotel")
		}
	case 3: // Bus
		var bus models.Bus
		if result := config.DB.First(&bus, "location = ?", reservation.ResvKey); result.Error != nil {
			return result.Error
		}
		available = bus.NumAvail
		if available > 0 {
			bus.NumAvail--
			config.DB.Save(&bus)
		} else {
			return fmt.Errorf("no available seats on bus")
		}
	default:
		return fmt.Errorf("invalid reservation type")
	}

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

func UpdateReservation(key string, updatedReservation models.Reservation) error {
	var reservation models.Reservation
	if result := config.DB.First(&reservation, "resv_key = ?", key); result.Error != nil {
		return result.Error
	}
	if result := config.DB.Model(&reservation).Updates(updatedReservation); result.Error != nil {
		return result.Error
	}
	fmt.Println("Reservation updated successfully")
	return nil
}

func DeleteReservation(key string) error {
	if result := config.DB.Delete(&models.Reservation{}, "resv_key = ?", key); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetReservationByKey(key string) (models.Reservation, error) {
	var reservation models.Reservation
	if result := config.DB.First(&reservation, "resv_key = ?", key); result.Error != nil {
		return models.Reservation{}, result.Error
	}
	return reservation, nil
}

func GetReservationsByType(resvType int) ([]models.Reservation, error) {
	var reservations []models.Reservation
	if result := config.DB.Find(&reservations, "resv_type = ?", resvType); result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}

func GetReservationsByCustomer(custName string) ([]models.Reservation, error) {
	var reservations []models.Reservation
	if result := config.DB.Find(&reservations, "cust_name = ?", custName); result.Error != nil {
		return nil, result.Error
	}
	return reservations, nil
}
