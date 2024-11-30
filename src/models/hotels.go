package models

import "gorm.io/gorm"

type Hotel struct {
	gorm.Model
	Location string `gorm:"primaryKey"`
	Price    int    `gorm:"not null"`
	NumRooms int    `gorm:"not null"`
	NumAvail int    `gorm:"not null"`
}
