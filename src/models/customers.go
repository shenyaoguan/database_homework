package models

type Customer struct {
	CustName string `gorm:"not null"`
	CustID   int    `gorm:"primaryKey"`
}
