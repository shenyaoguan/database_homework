package controllers

import (
	"src/config"
	"src/models"
	"testing"
)

func TestAddReservation(t *testing.T) {
	config.ConnectDatabase()
	reservation := models.Reservation{
		CustName: "TestCustomer",
		ResvType: 1,
		ResvKey:  "TestFlight",
	}
	err := AddReservation(reservation)
	if err != nil {
		t.Errorf("Failed to add reservation: %v", err)
	}
}

func TestGetAllReservations(t *testing.T) {
	config.ConnectDatabase()
	_, err := GetAllReservations()
	if err != nil {
		t.Errorf("Failed to get all reservations: %v", err)
	}
}

func TestUpdateReservation(t *testing.T) {
	config.ConnectDatabase()
	reservation := models.Reservation{
		CustName: "TestCustomer",
		ResvType: 2,
		ResvKey:  "TestHotel",
	}
	err := UpdateReservation("TestFlight", reservation)
	if err != nil {
		t.Errorf("Failed to update reservation: %v", err)
	}
}

func TestDeleteReservation(t *testing.T) {
	config.ConnectDatabase()
	err := DeleteReservation("TestFlight")
	if err != nil {
		t.Errorf("Failed to delete reservation: %v", err)
	}
}
