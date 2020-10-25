package models

import (
	"time"
)

// RecipeIngredients entity
type RecipeIngredients struct {
	ID           uint `gorm:"primary_key"`
	RecipeID     uint
	Recipe       Recipe
	IngredientID uint
	Ingredient   Ingredient
	Quantity     uint32

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
