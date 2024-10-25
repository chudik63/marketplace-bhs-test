package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id            int    `gorm:"uniqueIndex;not null"`
	Username      string `gorm:"not null"`
	Password_hash string `gorm:"not null"`
	Assets        []Asset
}
