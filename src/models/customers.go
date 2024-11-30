package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustName string `gorm:"primaryKey;not null"`
	CustID   int    `gorm:"unique;autoIncrement"`
}
