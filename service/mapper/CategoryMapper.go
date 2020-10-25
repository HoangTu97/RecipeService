package mapper

import (
	"Food/dto"
	"Food/models"
)

func ToCategoryDTO(entity models.Category) dto.CategoryDTO {
	return dto.CategoryDTO{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
		Name:      entity.Name,
		Image:     entity.Image,
	}
}

func ToCategory(dto dto.CategoryDTO) models.Category {
	return models.Category{
		ID:        dto.ID,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
		Name:      dto.Name,
		Image:     dto.Image,
	}
}

func ToCategoryDTOS(entityList []models.Category) []dto.CategoryDTO {
	dtos := make([]dto.CategoryDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = ToCategoryDTO(v)
	}

	return dtos
}

func ToCategories(dtoList []dto.CategoryDTO) []models.Category {
	entities := make([]models.Category, len(dtoList))

	for i, v := range dtoList {
		entities[i] = ToCategory(v)
	}

	return entities
}
