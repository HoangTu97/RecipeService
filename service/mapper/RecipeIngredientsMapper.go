package mapper

import (
	"Food/dto"
	"Food/models"
)

type RecipeIngredient interface {
	ToDTO(entity models.RecipeIngredients) dto.RecipeIngredientsDTO
	ToEntity(dto dto.RecipeIngredientsDTO) models.RecipeIngredients
	ToDTOS(entityList []models.RecipeIngredients) []dto.RecipeIngredientsDTO
	ToEntities(dtoList []dto.RecipeIngredientsDTO) []models.RecipeIngredients
}

type recipeIngredient struct {}

func NewRecipeIngredient() RecipeIngredient {
	return &recipeIngredient{}
}

func (m *recipeIngredient) ToDTO(entity models.RecipeIngredients) dto.RecipeIngredientsDTO {
	return dto.RecipeIngredientsDTO{
		ID:           entity.ID,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
		DeletedAt:    entity.DeletedAt,
		RecipeID:     entity.RecipeID,
		IngredientID: entity.IngredientID,
	}
}

func (m *recipeIngredient) ToEntity(dto dto.RecipeIngredientsDTO) models.RecipeIngredients {
	return models.RecipeIngredients{
		ID:           dto.ID,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
		DeletedAt:    dto.DeletedAt,
		RecipeID:     dto.RecipeID,
		IngredientID: dto.IngredientID,
	}
}

func (m *recipeIngredient) ToDTOS(entityList []models.RecipeIngredients) []dto.RecipeIngredientsDTO {
	dtos := make([]dto.RecipeIngredientsDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = m.ToDTO(v)
	}

	return dtos
}

func (m *recipeIngredient) ToEntities(dtoList []dto.RecipeIngredientsDTO) []models.RecipeIngredients {
	entities := make([]models.RecipeIngredients, len(dtoList))

	for i, v := range dtoList {
		entities[i] = m.ToEntity(v)
	}

	return entities
}
