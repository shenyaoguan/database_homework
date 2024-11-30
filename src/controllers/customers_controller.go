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

// Add more functions for updating and deleting customers
