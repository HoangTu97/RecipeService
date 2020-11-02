package repository

import (
	"Food/models"

	"gorm.io/gorm"
)

type recipe_ingredients struct {
	db *gorm.DB
}

type RecipeIngredients interface {
	FindAll() []models.RecipeIngredients
	FindOne(id uint) models.RecipeIngredients
	FindByRecipeIDs(recipeIDs []uint) []models.RecipeIngredients
	FindByRecipeID(recipeID uint) []models.RecipeIngredients
}

func NewRecipeIngredients(db *gorm.DB) RecipeIngredients {
	return &recipe_ingredients{db: db}
}

func (r *recipe_ingredients) FindAll() []models.RecipeIngredients {
	var recipeIngredientsLst []models.RecipeIngredients
	r.db.Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}

func (r *recipe_ingredients) FindOne(id uint) models.RecipeIngredients {
	var recipeIngredient models.RecipeIngredients
	r.db.First(&recipeIngredient, id)
	return recipeIngredient
}

func (r *recipe_ingredients) FindByRecipeIDs(recipeIDs []uint) []models.RecipeIngredients {
	var recipeIngredientsLst []models.RecipeIngredients
	r.db.Where("recipe_id IN (?)", recipeIDs).Preload("Ingredient").Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}

func (r *recipe_ingredients) FindByRecipeID(recipeID uint) []models.RecipeIngredients {
	var recipeIngredientsLst []models.RecipeIngredients
	r.db.Where("recipe_id = ?", recipeID).Preload("Ingredient").Find(&recipeIngredientsLst)
	return recipeIngredientsLst
}
