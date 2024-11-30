package models

type Flight struct {
	FlightNum string `gorm:"primaryKey"`
	Price     int    `gorm:"not null"`
	NumSeats  int    `gorm:"not null"`
	NumAvail  int    `gorm:"not null"`
	FromCity  string `gorm:"not null"`
	ArivCity  string `gorm:"not null"`
}
