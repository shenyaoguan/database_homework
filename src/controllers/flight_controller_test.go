package controllers

import (
	"src/config"
	"src/models"
	"testing"
)

func TestAddFlight(t *testing.T) {
	config.ConnectDatabase()
	flight := models.Flight{
		FlightNum: "TestFlight",
		Price:     200,
		NumSeats:  100,
		NumAvail:  100,
		FromCity:  "CityA",
		ArivCity:  "CityB",
	}
	err := AddFlight(flight)
	if err != nil {
		t.Errorf("Failed to add flight: %v", err)
	}
}

func TestGetAllFlights(t *testing.T) {
	config.ConnectDatabase()
	_, err := GetAllFlights()
	if err != nil {
		t.Errorf("Failed to get all flights: %v", err)
	}
}

func TestUpdateFlight(t *testing.T) {
	config.ConnectDatabase()
	flight := models.Flight{
		FlightNum: "TestFlight",
		Price:     250,
		NumSeats:  90,
		NumAvail:  90,
		FromCity:  "CityA",
		ArivCity:  "CityC",
	}
	err := UpdateFlight("TestFlight", flight)
	if err != nil {
		t.Errorf("Failed to update flight: %v", err)
	}
}

func TestDeleteFlight(t *testing.T) {
	config.ConnectDatabase()
	err := DeleteFlight("TestFlight")
	if err != nil {
		t.Errorf("Failed to delete flight: %v", err)
	}
}
