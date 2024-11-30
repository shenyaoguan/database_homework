package models

type Hotel struct {
	Location string `gorm:"primaryKey"`
	Price    int    `gorm:"not null"`
	NumRooms int    `gorm:"not null"`
	NumAvail int    `gorm:"not null"`
}
