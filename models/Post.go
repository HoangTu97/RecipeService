package models

import (
	"time"
)

// Post entity
type Post struct {
	ID uint `gorm:"primary_key"`

	Photo       string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:text"`
	Type        string `gorm:"type:varchar(255)"`
	HashTags    string `gorm:"type:varchar(255)"`

	UserID   uint
	User     User
	RecipeID uint
	Recipe   Recipe
	Comments []Comment

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
