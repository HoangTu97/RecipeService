package mapper

import (
	"Food/dto"
	"Food/models"
)

type Category interface {
	ToDTO(entity models.Category) dto.CategoryDTO
	ToEntity(dto dto.CategoryDTO) models.Category
	ToDTOS(entityList []models.Category) []dto.CategoryDTO
	ToEntities(dtoList []dto.CategoryDTO) []models.Category
}

type category struct {}

func NewCategory() Category {
	return &category{}
}

func (m *category) ToDTO(entity models.Category) dto.CategoryDTO {
	return dto.CategoryDTO{
		ID:        entity.ID,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
		Name:      entity.Name,
		Image:     entity.Image,
	}
}

func (m *category) ToEntity(dto dto.CategoryDTO) models.Category {
	return models.Category{
		ID:        dto.ID,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
		DeletedAt: dto.DeletedAt,
		Name:      dto.Name,
		Image:     dto.Image,
	}
}

func (m *category) ToDTOS(entityList []models.Category) []dto.CategoryDTO {
	dtos := make([]dto.CategoryDTO, len(entityList))

	for i, v := range entityList {
		dtos[i] = m.ToDTO(v)
	}

	return dtos
}

func (m *category) ToEntities(dtoList []dto.CategoryDTO) []models.Category {
	entities := make([]models.Category, len(dtoList))

	for i, v := range dtoList {
		entities[i] = m.ToEntity(v)
	}

	return entities
}
