package service

import (
	"Food/models"
	"Food/repository"
)

func FindRecipeIngredientsByRecipeIDs(recipeIDs []uint) []models.RecipeIngredients {
	return repository.FindRecipeIngredientsByRecipeIDs(recipeIDs)
}

func FindRecipeIngredientsByRecipeID(recipeID uint) []models.RecipeIngredients {
	return repository.FindRecipeIngredientsByRecipeID(recipeID)
}
