package mapper

import (
	"Food/dto"
	"Food/models"
)

func ToCommentDTO(entity models.Comment) dto.CommentDTO {
	return dto.CommentDTO{
		ID:          entity.ID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
		Description: entity.Description,
		UserID:      entity.UserID,
		PostID:      entity.PostID,
	}
}

func ToComment(dto dto.CommentDTO) models.Comment {
	return models.Comment{
		ID:          dto.ID,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		DeletedAt:   dto.DeletedAt,
		Description: dto.Description,
		UserID:      dto.UserID,
		PostID:      dto.PostID,
	}
}

func ToCommentDTOS(entityList []models.Comment) []dto.CommentDTO {
	dtos := make([]dto.CommentDTO, len(entityList))

	for i, entity := range entityList {
		dtos[i] = ToCommentDTO(entity)
	}

	return dtos
}

func ToComments(dtoList []dto.CommentDTO) []models.Comment {
	entities := make([]models.Comment, len(dtoList))

	for i, dto := range dtoList {
		entities[i] = ToComment(dto)
	}

	return entities
}
