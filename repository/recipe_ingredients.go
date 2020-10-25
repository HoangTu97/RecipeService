package repository

import (
	"Food/config"
	"Food/models"
)

func FindRecipeIngredientsAll() []models.RecipeIngredients {
	var recipeIngredientsLst []models.RecipeIngredients
	config.GetDB().Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}

func FindOneRecipeIngredients(id uint) models.RecipeIngredients {
	var recipeIngredient models.RecipeIngredients
	config.GetDB().First(&recipeIngredient, id)
	return recipeIngredient
}

func FindRecipeIngredientsByRecipeIDs(recipeIDs []uint) []models.RecipeIngredients {
	var recipeIngredientsLst []models.RecipeIngredients
	config.GetDB().Where("recipe_id IN (?)", recipeIDs).Preload("Ingredient").Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}

func FindRecipeIngredientsByRecipeID(recipeID uint) []models.RecipeIngredients {
	var recipeIngredientsLst []models.RecipeIngredients
	config.GetDB().Where("recipe_id = ?", recipeID).Preload("Ingredient").Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}
