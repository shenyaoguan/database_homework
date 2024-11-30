package controllers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"src/config"
	"src/models"
	"testing"
)

func TestAddCustomer(t *testing.T) {
	config.ConnectDatabase()
	customer := models.Customer{
		CustName: "TestCustomer",
		CustID:   1,
	}
	err := AddCustomer(customer)
	if err != nil {
		t.Errorf("Failed to add customer: %v", err)
	}
}

func TestGetAllCustomers(t *testing.T) {
	config.ConnectDatabase()
	_, err := GetAllCustomers()
	if err != nil {
		t.Errorf("Failed to get all customers: %v", err)
	}
}

func TestUpdateCustomer(t *testing.T) {
	config.ConnectDatabase()
	customer := models.Customer{
		CustName: "UpdatedCustomer",
		CustID:   1,
	}
	err := UpdateCustomer("TestCustomer", customer)
	if err != nil {
		t.Errorf("Failed to update customer: %v", err)
	}
}

func TestDeleteCustomer(t *testing.T) {
	config.ConnectDatabase()
	err := DeleteCustomer("UpdatedCustomer")
	if err != nil {
		t.Errorf("Failed to delete customer: %v", err)
	}
}

func TestGetCustomerItinerary(t *testing.T) {
	config.ConnectDatabase()

	// Add test data
	customer := models.Customer{CustName: "TestCustomer", CustID: 200}
	config.DB.Create(&customer)

	flight := models.Flight{FlightNum: "FL123", NumSeats: 100, NumAvail: 100}
	config.DB.Create(&flight)

	hotel := models.Hotel{Location: "TestLocationHotel", NumRooms: 50, NumAvail: 50, Price: 100}
	config.DB.Create(&hotel)

	bus := models.Bus{Location: "TestLocationBus", NumBus: 10, NumAvail: 10}
	config.DB.Create(&bus)

	reservation1 := models.Reservation{CustName: "TestCustomer", ResvType: 1, ResvKey: "FL123"}
	reservation2 := models.Reservation{CustName: "TestCustomer", ResvType: 2, ResvKey: "TestLocationHotel"}
	reservation3 := models.Reservation{CustName: "TestCustomer", ResvType: 3, ResvKey: "TestLocationBus"}
	config.DB.Create(&reservation1)
	config.DB.Create(&reservation2)
	config.DB.Create(&reservation3)

	// Test GetCustomerItinerary
	itinerary, err := GetCustomerItinerary("TestCustomer")
	if err != nil {
		t.Errorf("Failed to get customer itinerary: %v", err)
	}

	if len(itinerary) != 3 {
		t.Errorf("Expected 3 itinerary items, got %d", len(itinerary))
	}

	// Clean up test data
	config.DB.Delete(&reservation1)
	config.DB.Delete(&reservation2)
	config.DB.Delete(&reservation3)
	config.DB.Delete(&flight)
	config.DB.Delete(&hotel)
	config.DB.Delete(&bus)
	config.DB.Delete(&customer)
}

func TestCheckReservationIntegrity(t *testing.T) {
	config.ConnectDatabase()
	db := config.DB
	// Insert mock data for flights, hotels, buses, and customers
	db.Create(&models.Flight{FlightNum: "F123", Price: 500, NumSeats: 100, NumAvail: 50, FromCity: "CityA", ArivCity: "CityB"})
	db.Create(&models.Hotel{Location: "CityA", Price: 300, NumRooms: 50, NumAvail: 20})
	db.Create(&models.Bus{Location: "CityB", Price: 100, NumBus: 10, NumAvail: 5})
	db.Create(&models.Customer{CustName: "John Doe", CustID: 1})

	// Insert valid reservations for the customer
	db.Create(&models.Reservation{CustName: "John Doe", ResvType: 1, ResvKey: "F123"})  // Flight reservation
	db.Create(&models.Reservation{CustName: "John Doe", ResvType: 2, ResvKey: "CityA"}) // Hotel reservation
	db.Create(&models.Reservation{CustName: "John Doe", ResvType: 3, ResvKey: "CityB"}) // Bus reservation

	// Test Case 1: Valid reservations
	err := CheckReservationIntegrity("John Doe")
	assert.NoError(t, err, "Expected no error for valid reservations")

	// Test Case 2: Insufficient availability for a flight (ResvKey refers to F123)
	db.Model(&models.Flight{}).Where("flight_num = ?", "F123").Update("num_avail", 0)
	err = CheckReservationIntegrity("John Doe")
	assert.EqualError(t, err, "flight reservation has insufficient availability", "Expected error for insufficient flight availability")

	// Restore flight availability
	db.Model(&models.Flight{}).Where("flight_num = ?", "F123").Update("num_avail", 50)

	// Test Case 3: Invalid hotel reservation (ResvKey refers to an invalid location)
	db.Delete(&models.Hotel{}, "location = ?", "CityA") // Remove hotel record
	err = CheckReservationIntegrity("John Doe")
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound), "Expected error for invalid hotel reservation")

	// Restore hotel record
	db.Create(&models.Hotel{Location: "CityA", Price: 300, NumRooms: 50, NumAvail: 20})

	// Test Case 4: Invalid bus reservation (ResvKey refers to an invalid location)
	db.Delete(&models.Bus{}, "location = ?", "CityB") // Remove bus record
	err = CheckReservationIntegrity("John Doe")
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound), "Expected error for invalid bus reservation")

	// Restore bus record
	db.Create(&models.Bus{Location: "CityB", Price: 100, NumBus: 10, NumAvail: 5})

	// Test Case 5: Unknown reservation type (ResvKey refers to an unknown type)
	db.Create(&models.Reservation{CustName: "John Doe", ResvType: 4, ResvKey: "Unknown"})
	err = CheckReservationIntegrity("John Doe")
	assert.EqualError(t, err, "unknown reservation type: 4", "Expected error for unknown reservation type")

	// Remove unknown reservation
	db.Delete(&models.Reservation{}, "cust_name = ? AND resv_type = ? AND resv_key = ?", "John Doe", 4, "Unknown")

	// Test Case 6: Restore and validate integrity after fixing all issues
	// Restore bus and hotel
	db.Create(&models.Hotel{Location: "CityA", Price: 300, NumRooms: 50, NumAvail: 20})
	db.Create(&models.Bus{Location: "CityB", Price: 100, NumBus: 10, NumAvail: 5})
	err = CheckReservationIntegrity("John Doe")
	assert.NoError(t, err, "Expected no error after restoring all resources")

	// Clean up
	db.Delete(&models.Reservation{}, "cust_name = ?", "John Doe")
	db.Delete(&models.Flight{}, "flight_num = ?", "F123")
	db.Delete(&models.Hotel{}, "location = ?", "CityA")
	db.Delete(&models.Bus{}, "location = ?", "CityB")
	db.Delete(&models.Customer{}, "cust_name = ?", "John Doe")

}
