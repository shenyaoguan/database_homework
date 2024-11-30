package models

type Bus struct {
	Location string `gorm:"primaryKey"`
	Price    int    `gorm:"not null"`
	NumBus   int    `gorm:"not null"`
	NumAvail int    `gorm:"not null"`
}
