package models

import (
	"time"
)

// Ingredient entity
type Ingredient struct {
	ID                uint     `gorm:"primary_key"`
	Name              string   `gorm:"type:varchar(255)"`
	Image             string   `gorm:"type:varchar(255)"`
	Description       string   `gorm:"type:text"`
	RecipeIngredients []RecipeIngredients

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
