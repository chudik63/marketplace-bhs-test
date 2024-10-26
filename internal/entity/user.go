package entity

import "gorm.io/gorm"

// User godoc
// @Description Information about a user in the marketplace
type User struct {
	gorm.Model
	// Username chosen by the user, must be unique
	// Required: true
	Username string `gorm:"uniqueIndex;not null" json:"username" example:"john_doe"`

	// Hashed password for user authentication
	// Required: true
	Password_hash string `gorm:"not null" json:"password_hash" example:"hashed_password_value"`

	// User's current balance
	Balance float64 `gorm:"not null" json:"balance" example:"100.50"`
}
