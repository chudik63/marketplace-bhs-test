package entity

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Name        string  `gorm:"uniqueIndex;not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	UserID      uint    `gorm:"not null"`
}
