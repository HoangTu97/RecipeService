package service

import (
	"Food/models"
	"Food/repository"
)

type RecipeIngredients interface {
	FindByRecipeIDs(recipeIDs []uint) []models.RecipeIngredients
	FindByRecipeID(recipeID uint) []models.RecipeIngredients
}

type recipe_ingredients struct {
	repository repository.RecipeIngredients
}

func NewRecipeIngredients(repository repository.RecipeIngredients) RecipeIngredients {
	return &recipe_ingredients{repository: repository}
}

func (s *recipe_ingredients) FindByRecipeIDs(recipeIDs []uint) []models.RecipeIngredients {
	return s.repository.FindByRecipeIDs(recipeIDs)
}

func (s *recipe_ingredients) FindByRecipeID(recipeID uint) []models.RecipeIngredients {
	return s.repository.FindByRecipeID(recipeID)
}
