package mapper

import (
	"Food/dto"
	"Food/models"
)

func ToRecipeIngredientsDTO(entity models.RecipeIngredients) dto.RecipeIngredientsDTO {
	return dto.RecipeIngredientsDTO{
		ID:           entity.ID,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
		DeletedAt:    entity.DeletedAt,
		RecipeID:     entity.RecipeID,
		IngredientID: entity.IngredientID,
	}
}

func ToRecipeIngredient(dto dto.RecipeIngredientsDTO) models.RecipeIngredients {
	return models.RecipeIngredients{
		ID:           dto.ID,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
		DeletedAt:    dto.DeletedAt,
		RecipeID:     dto.RecipeID,
		IngredientID: dto.IngredientID,
	}
}

func ToRecipeIngredientsDTOS(entityList []models.RecipeIngredients) []dto.RecipeIngredientsDTO {
	dtos := make([]dto.RecipeIngredientsDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = ToRecipeIngredientsDTO(v)
	}

	return dtos
}

func ToRecipeIngredients(dtoList []dto.RecipeIngredientsDTO) []models.RecipeIngredients {
	entities := make([]models.RecipeIngredients, len(dtoList))

	for i, v := range dtoList {
		entities[i] = ToRecipeIngredient(v)
	}

	return entities
}
