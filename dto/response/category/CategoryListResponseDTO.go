package category

import (
	"Food/helpers/page"
	"Food/models"
)

// CategoryListItemResponseDTO
type CategoryListItemResponseDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Image         string `json:"image"`
	NumberRecipes int    `json:"numberRecipes"`
}

// CategoryListResponseDTO godoc
type CategoryListResponseDTO struct {
	Items      []CategoryListItemResponseDTO `json:"items"`
	TotalItems int64                         `json:"totalItems"`
}

// CreateCategoryListResponseDTOFromPage create page from page models.Category
func CreateCategoryListResponseDTOFromPage(page page.Page) *CategoryListResponseDTO {
	result := make([]CategoryListItemResponseDTO, page.GetTotalElements())
	for i, v := range page.GetContent() {
		entity := v.(models.Category)
		recipes := entity.Recipes

		result[i] = CategoryListItemResponseDTO{
			ID:            entity.ID,
			Name:          entity.Name,
			Image:         entity.Image,
			NumberRecipes: len(recipes),
		}
	}
	return &CategoryListResponseDTO{
		Items:      result,
		TotalItems: page.GetTotalItems(),
	}
}
