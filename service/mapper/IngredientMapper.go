package mapper

import (
	"Food/dto"
	"Food/models"
)

type Ingredient interface {
	ToDTO(entity models.Ingredient) dto.IngredientDTO
	ToEntity(dto dto.IngredientDTO) models.Ingredient
	ToDTOS(entityList []models.Ingredient) []dto.IngredientDTO
	ToEntities(dtoList []dto.IngredientDTO) []models.Ingredient
}

type ingredient struct {}

func NewIngredient() Ingredient {
	return &ingredient{}
}

func (m *ingredient) ToDTO(entity models.Ingredient) dto.IngredientDTO {
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

func (m *ingredient) ToEntity(dto dto.IngredientDTO) models.Ingredient {
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

func (m *ingredient) ToDTOS(entityList []models.Ingredient) []dto.IngredientDTO {
	dtos := make([]dto.IngredientDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = m.ToDTO(v)
	}

	return dtos
}

func (m *ingredient) ToEntities(dtoList []dto.IngredientDTO) []models.Ingredient {
	entities := make([]models.Ingredient, len(dtoList))

	for i, v := range dtoList {
		entities[i] = m.ToEntity(v)
	}

	return entities
}
