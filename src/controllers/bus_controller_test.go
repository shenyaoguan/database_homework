package controllers

import (
	"src/config"
	"src/models"
	"testing"
)

func TestAddBus(t *testing.T) {
	config.ConnectDatabase()
	bus := models.Bus{
		Location: "TestLocation",
		Price:    50,
		NumBus:   10,
		NumAvail: 10,
	}
	err := AddBus(bus)
	if err != nil {
		t.Errorf("Failed to add bus: %v", err)
	}
}

func TestGetAllBuses(t *testing.T) {
	config.ConnectDatabase()
	_, err := GetAllBuses()
	if err != nil {
		t.Errorf("Failed to get all buses: %v", err)
	}
}

func TestUpdateBus(t *testing.T) {
	config.ConnectDatabase()
	bus := models.Bus{
		Location: "TestLocation",
		Price:    60,
		NumBus:   8,
		NumAvail: 8,
	}
	err := UpdateBus("TestLocation", bus)
	if err != nil {
		t.Errorf("Failed to update bus: %v", err)
	}
}

func TestDeleteBus(t *testing.T) {
	config.ConnectDatabase()
	err := DeleteBus("TestLocation")
	if err != nil {
		t.Errorf("Failed to delete bus: %v", err)
	}
}
