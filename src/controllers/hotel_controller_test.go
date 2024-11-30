package controllers

import (
	"src/config"
	"src/models"
	"testing"
)

func TestAddHotel(t *testing.T) {
	config.ConnectDatabase()
	hotel := models.Hotel{
		Location: "TestLocation",
		Price:    100,
		NumRooms: 20,
		NumAvail: 20,
	}
	err := AddHotel(hotel)
	if err != nil {
		t.Errorf("Failed to add hotel: %v", err)
	}
}

func TestGetAllHotels(t *testing.T) {
	config.ConnectDatabase()
	_, err := GetAllHotels()
	if err != nil {
		t.Errorf("Failed to get all hotels: %v", err)
	}
}

func TestUpdateHotel(t *testing.T) {
	config.ConnectDatabase()
	hotel := models.Hotel{
		Location: "TestLocation",
		Price:    120,
		NumRooms: 18,
		NumAvail: 18,
	}
	err := UpdateHotel("TestLocation", hotel)
	if err != nil {
		t.Errorf("Failed to update hotel: %v", err)
	}
}

func TestDeleteHotel(t *testing.T) {
	config.ConnectDatabase()
	err := DeleteHotel("TestLocation")
	if err != nil {
		t.Errorf("Failed to delete hotel: %v", err)
	}
}
