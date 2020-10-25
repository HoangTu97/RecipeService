package repository

import (
	"Food/config"
	"Food/models"
)

func FindAllIngredient() []models.Ingredient {
	var ingredients []models.Ingredient
	config.GetDB().Find(&ingredients)
	return ingredients
}

func FindOneIngredient(id uint) (models.Ingredient, error) {
	var ingredient models.Ingredient
	result := config.GetDB().First(&ingredient, id)
	if result.Error != nil {
		return models.Ingredient{}, result.Error
	}
	return ingredient, nil
}

func FindIngredientByName(name string) []models.Ingredient {
	var ingredients []models.Ingredient
	config.GetDB().Where("name LIKE ?", "%" + name + "%").Find(&ingredients)
	return ingredients
}
