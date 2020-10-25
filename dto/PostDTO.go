package dto

import (
	"time"
)

// PostDTO godoc
type PostDTO struct {
	ID uint

	Photo       string
	Description string
	Type        string
	HashTags    []string

	UserID   uint
	RecipeID uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
