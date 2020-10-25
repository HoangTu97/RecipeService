package dto

import (
	"time"
)

// CommentDTO godoc
type CommentDTO struct {
	ID uint

	Description string 

	UserID uint
	PostID uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
