package controllers

import (
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
	err := DeleteCustomer("TestCustomer")
	if err != nil {
		t.Errorf("Failed to delete customer: %v", err)
	}
}
