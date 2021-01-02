package dto

import (
	"time"
)

// RecipeDTO godoc
type RecipeDTO struct {
	ID          uint
	Name        string
	Description string
	Image       string
	Photos      []string
	Duration    uint32
	Steps       []string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
