package mapper

import (
	"Food/dto"
	"Food/models"
)

func ToIngredientDTO(entity models.Ingredient) dto.IngredientDTO {
	return dto.IngredientDTO{
		ID:          entity.ID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
		Name:        entity.Name,
		Image:       entity.Image,
		Description: entity.Description,
	}
}

func ToIngredient(dto dto.IngredientDTO) models.Ingredient {
	return models.Ingredient{
		ID:          dto.ID,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		DeletedAt:   dto.DeletedAt,
		Name:        dto.Name,
		Image:       dto.Image,
		Description: dto.Description,
	}
}

func ToIngredientDTOS(entityList []models.Ingredient) []dto.IngredientDTO {
	dtos := make([]dto.IngredientDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = ToIngredientDTO(v)
	}

	return dtos
}

func ToIngredients(dtoList []dto.IngredientDTO) []models.Ingredient {
	entities := make([]models.Ingredient, len(dtoList))

	for i, v := range dtoList {
		entities[i] = ToIngredient(v)
	}

	return entities
}
