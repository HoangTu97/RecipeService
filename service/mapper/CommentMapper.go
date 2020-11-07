package mapper

import (
	"Food/dto"
	"Food/models"
)

type Comment interface {
	ToDTO(entity models.Comment) dto.CommentDTO
	ToEntity(dto dto.CommentDTO) models.Comment
	ToDTOS(entityList []models.Comment) []dto.CommentDTO
	ToEntities(dtoList []dto.CommentDTO) []models.Comment
}

type comment struct {}

func NewComment() Comment {
	return &comment{}
}

func (m *comment) ToDTO(entity models.Comment) dto.CommentDTO {
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

func (m *comment) ToEntity(dto dto.CommentDTO) models.Comment {
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

func (m *comment) ToDTOS(entityList []models.Comment) []dto.CommentDTO {
	dtos := make([]dto.CommentDTO, len(entityList))

	for i, entity := range entityList {
		dtos[i] = m.ToDTO(entity)
	}

	return dtos
}

func (m *comment) ToEntities(dtoList []dto.CommentDTO) []models.Comment {
	entities := make([]models.Comment, len(dtoList))

	for i, dto := range dtoList {
		entities[i] = m.ToEntity(dto)
	}

	return entities
}
