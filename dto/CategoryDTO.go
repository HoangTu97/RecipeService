package dto

import (
	"time"
)

// CategoryDTO godoc
type CategoryDTO struct {
	ID    uint
	Name  string
	Image string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
