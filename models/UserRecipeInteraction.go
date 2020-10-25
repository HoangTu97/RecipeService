package models

import (
	"time"
)

// UserRecipeInteraction entity
type UserRecipeInteraction struct {
	ID       uint   `gorm:"primary_key"`
	UserID   uint   `gorm:"index:search_key"`
	RecipeID uint   `gorm:"index:search_key"`
	Rating   int    ``
	Review   string `gorm:"type:text"`

	CreatedAt time.Time
}
