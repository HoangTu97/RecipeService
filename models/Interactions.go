package models

import (
	"gorm.io/gorm"
)

// Interactions entity
type Interactions struct {
	gorm.Model
	UserID     uint
	LocationID uint
	Value      int32
}
