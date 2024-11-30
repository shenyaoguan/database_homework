package models

import "gorm.io/gorm"

type Bus struct {
	gorm.Model
	Location string `gorm:"primaryKey"`
	Price    int    `gorm:"not null"`
	NumBus   int    `gorm:"not null"`
	NumAvail int    `gorm:"not null"`
}
