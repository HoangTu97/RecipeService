package category

import (
	"Food/dto"
)

// CategoryMiniListItemResponseDTO godoc
type CategoryMiniListItemResponseDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// CategoryMiniListResponseDTO godoc
type CategoryMiniListResponseDTO struct {
	Items      []CategoryMiniListItemResponseDTO `json:"items"`
	TotalItems int64                             `json:"totalItems"`
}

// CreateCategoryMiniListResponseDTO create page from page models.Category
func CreateCategoryMiniListResponseDTO(dtos []dto.CategoryDTO) *CategoryMiniListResponseDTO {
	result := make([]CategoryMiniListItemResponseDTO, len(dtos))
	for i, v := range dtos {
		entity := v

		result[i] = CategoryMiniListItemResponseDTO{
			ID:    entity.ID,
			Name:  entity.Name,
			Image: entity.Image,
		}
	}
	return &CategoryMiniListResponseDTO{
		Items:      result,
		TotalItems: int64(len(result)),
	}
}
