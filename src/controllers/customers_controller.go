package controllers

import (
	"fmt"
	"src/config"
	"src/models"
)

func AddCustomer(customer models.Customer) error {
	if result := config.DB.Create(&customer); result.Error != nil {
		return result.Error
	}
	fmt.Println("Customer added successfully")
	return nil
}

func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	if result := config.DB.Find(&customers); result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func GetCustomerByName(name string) (models.Customer, error) {
	var customer models.Customer
	if result := config.DB.First(&customer, "cust_name = ?", name); result.Error != nil {
		return models.Customer{}, result.Error
	}
	return customer, nil
}

// Add more functions for updating and deleting customers
func UpdateCustomer(name string, updatedCustomer models.Customer) error {
	var customer models.Customer
	if result := config.DB.First(&customer, "cust_name = ?", name); result.Error != nil {
		return result.Error
	}
	if result := config.DB.Model(&customer).Updates(updatedCustomer); result.Error != nil {
		return result.Error
	}
	fmt.Println("Customer updated successfully")
	return nil
}

func DeleteCustomer(name string) error {
	if result := config.DB.Delete(&models.Customer{}, "cust_name = ?", name); result.Error != nil {
		return result.Error
	}
	fmt.Println("Customer deleted successfully")
	return nil
}

func GetCustomerItinerary(custName string) ([]interface{}, error) {
	var reservations []models.Reservation
	if result := config.DB.Find(&reservations, "cust_name = ?", custName); result.Error != nil {
		return nil, result.Error
	}

	var itinerary []interface{}
	for _, res := range reservations {
		switch res.ResvType {
		case 1: // Flight
			var flight models.Flight
			if result := config.DB.First(&flight, "flight_num = ?", res.ResvKey); result.Error == nil {
				itinerary = append(itinerary, flight)
			}
		case 2: // Hotel
			var hotel models.Hotel
			if result := config.DB.First(&hotel, "location = ?", res.ResvKey); result.Error == nil {
				itinerary = append(itinerary, hotel)
			}
		case 3: // Bus
			var bus models.Bus
			if result := config.DB.First(&bus, "location = ?", res.ResvKey); result.Error == nil {
				itinerary = append(itinerary, bus)
			}
		}
	}

	return itinerary, nil
}
