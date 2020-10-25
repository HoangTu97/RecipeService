package dto

import (
	"time"
)

// IngredientDTO godoc
type IngredientDTO struct {
	ID          uint
	Name        string
	Image       string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
