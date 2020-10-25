package models

import (
	"time"
)

// Category entity
type Category struct {
	ID      uint     `gorm:"primary_key"`
	Name    string   `gorm:"type:varchar(255)"`
	Image   string   `gorm:"type:varchar(255)"`
	Recipes []Recipe `gorm:"many2many:cate_recipes;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
