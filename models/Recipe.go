package models

import (
	"time"
)

// Recipe entity
type Recipe struct {
	ID                uint   `gorm:"primary_key"`
	Name              string `gorm:"type:varchar(255)"`
	Description       string `gorm:"type:text"`
	Image             string `gorm:"type:varchar(255)"`
	Photos            string `gorm:"type:text"`
	Duration          uint32
	Categories        []Category   `gorm:"many2many:cate_recipes;"`
	RecipeIngredients []RecipeIngredients

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
