package config

import (
	"src/models"
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	ConnectDatabase()

	if DB == nil {
		t.Fatal("Database connection is nil")
	}

	if err := DB.Exec("SELECT 1").Error; err != nil {
		t.Fatalf("Failed to execute test query: %v", err)
	}
}

func TestAutoMigrate(t *testing.T) {
	ConnectDatabase()

	models := []interface{}{
		&models.Flight{},
		&models.Hotel{},
		&models.Bus{},
		&models.Customer{},
		&models.Reservation{},
	}

	for _, model := range models {
		if !DB.Migrator().HasTable(model) {
			t.Fatalf("Table for model %T does not exist", model)
		}
	}
}
