package models

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	CustName string `gorm:"not null"`
	ResvType int    `gorm:"not null"` // 1: Flight, 2: Hotel, 3: Bus
	ResvKey  string `gorm:"not null; primaryKey"`
}
