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

func GetCustomerByEmail(email string) (models.Customer, error) {
	var customer models.Customer
	if result := config.DB.First(&customer, "email = ?", email); result.Error != nil {
		return models.Customer{}, result.Error
	}
	return customer, nil
}

func GetCustomerByPhone(phone string) (models.Customer, error) {
	var customer models.Customer
	if result := config.DB.First(&customer, "phone = ?", phone); result.Error != nil {
		return models.Customer{}, result.Error
	}
	return customer, nil
}

// Add more functions for updating and deleting customers
func UpdateCustomer(email string, updatedCustomer models.Customer) error {
	var customer models.Customer
	if result := config.DB.First(&customer, "email = ?", email); result.Error != nil {
		return result.Error
	}
	if result := config.DB.Model(&customer).Updates(updatedCustomer); result.Error != nil {
		return result.Error
	}
	fmt.Println("Customer updated successfully")
	return nil
}

func DeleteCustomer(email string) error {
	if result := config.DB.Delete(&models.Customer{}, "email = ?", email); result.Error != nil {
		return result.Error
	}
	fmt.Println("Customer deleted successfully")
	return nil
}
