package entity

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Id          int     `gorm:"uniqueIndex;not null"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
}
