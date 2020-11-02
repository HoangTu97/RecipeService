package repository

import (
	"Food/models"

	"gorm.io/gorm"
)

type ingredient struct {
	db *gorm.DB
}

type Ingredient interface {
	FindAll() []models.Ingredient
	FindOne(id uint) (models.Ingredient, error)
	FindByName(name string) []models.Ingredient
}

func NewIngredient(db *gorm.DB) Ingredient {
	return &ingredient{db: db}
}

func (r *ingredient) FindAll() []models.Ingredient {
	var ingredients []models.Ingredient
	r.db.Find(&ingredients)
	return ingredients
}

func (r *ingredient) FindOne(id uint) (models.Ingredient, error) {
	var ingredient models.Ingredient
	result := r.db.First(&ingredient, id)
	if result.Error != nil {
		return models.Ingredient{}, result.Error
	}
	return ingredient, nil
}

func (r *ingredient) FindByName(name string) []models.Ingredient {
	var ingredients []models.Ingredient
	r.db.Where("name LIKE ?", "%" + name + "%").Find(&ingredients)
	return ingredients
}
