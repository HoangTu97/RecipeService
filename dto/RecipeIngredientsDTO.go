package dto

import (
	"time"
)

// RecipeIngredientsDTO godoc
type RecipeIngredientsDTO struct {
	ID           uint
	RecipeID     uint
	IngredientID uint
	Quantity     uint32

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
