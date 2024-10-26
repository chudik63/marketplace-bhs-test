package entity

import "gorm.io/gorm"

// Asset godoc
// @Description Information about an asset in the marketplace
type Asset struct {
	gorm.Model
	// Unique name for the asset
	// Required: true
	Name string `gorm:"uniqueIndex;not null" json:"name" example:"Artwork"`

	// Detailed description of the asset
	// Required: true
	Description string `gorm:"not null" json:"description" example:"Digital artwork"`

	// Price of the asset in the marketplace
	// Required: true
	Price float64 `gorm:"not null" json:"price" example:"100.50"`

	UserID uint64 `gorm:"not null" json:"user_id" example:"1"`
}
