package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string  `gorm:"uniqueIndex;not null"`
	Password_hash string  `gorm:"not null"`
	Assets        []Asset `gorm:"foreignKey:UserID"`
}
