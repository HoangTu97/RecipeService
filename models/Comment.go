package models

import (
	"time"
)

// Comment entity
type Comment struct {
	ID uint `gorm:"primary_key"`

	Description string `gorm:"type:text"`

	UserID uint
	User   User
	PostID uint
	Post   Post

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
