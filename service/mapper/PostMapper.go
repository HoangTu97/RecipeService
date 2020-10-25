package mapper

import (
	"Food/dto"
	"Food/helpers/converter"
	"Food/models"
)

func ToPostDTO(entity models.Post) dto.PostDTO {
	return dto.PostDTO{
		ID:          entity.ID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
		Photo:       entity.Photo,
		Description: entity.Description,
		Type:        entity.Type,
		HashTags:    converter.MustArrStr(entity.HashTags),
		UserID:      entity.UserID,
		RecipeID:    entity.RecipeID,
	}
}

func ToPost(dto dto.PostDTO) models.Post {
	return models.Post{
		ID:          dto.ID,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		DeletedAt:   dto.DeletedAt,
		Photo:       dto.Photo,
		Description: dto.Description,
		Type:        dto.Type,
		HashTags:    converter.ToStr(dto.HashTags),
		UserID:      dto.UserID,
		RecipeID:    dto.RecipeID,
	}
}

func ToPostDTOS(entityList []models.Post) []dto.PostDTO {
	dtos := make([]dto.PostDTO, len(entityList))

	for i, entity := range entityList {
		dtos[i] = ToPostDTO(entity)
	}

	return dtos
}

func ToPosts(dtoList []dto.PostDTO) []models.Post {
	entities := make([]models.Post, len(dtoList))

	for i, dto := range dtoList {
		entities[i] = ToPost(dto)
	}

	return entities
}
