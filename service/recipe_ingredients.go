package service

import (
	"Food/models"
	"Food/repository"
	"Food/service/mapper"
)

type RecipeIngredients interface {
	FindByRecipeIDs(recipeIDs []uint) []models.RecipeIngredients
	FindByRecipeID(recipeID uint) []models.RecipeIngredients
}

type recipe_ingredients struct {
	repository repository.RecipeIngredients
	mapper mapper.RecipeIngredient
}

func NewRecipeIngredients(repository repository.RecipeIngredients, mapper mapper.RecipeIngredient) RecipeIngredients {
	return &recipe_ingredients{repository: repository, mapper: mapper}
}

func (s *recipe_ingredients) FindByRecipeIDs(recipeIDs []uint) []models.RecipeIngredients {
	return s.repository.FindByRecipeIDs(recipeIDs)
}

func (s *recipe_ingredients) FindByRecipeID(recipeID uint) []models.RecipeIngredients {
	return s.repository.FindByRecipeID(recipeID)
}
