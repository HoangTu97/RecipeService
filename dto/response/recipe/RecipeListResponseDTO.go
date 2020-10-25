package recipe

import (
	"Food/models"
	"Food/helpers/page"
)

// RecipeListResponseDTO godoc
type RecipeListResponseDTO struct {
	Items      []RecipeListItemResponseDTO `json:"items"`
	TotalItems int64                       `json:"totalItems"`
}

// CreateRecipeListResponseDTOFromPage create page from page models.Recipe
func CreateRecipeListResponseDTOFromPage(page page.Page) *RecipeListResponseDTO {
	result := make([]RecipeListItemResponseDTO, page.GetTotalElements())
	for i, v := range page.GetContent() {
		entity := v.(models.Recipe)
		categories := entity.Categories

		cates := make([]string, len(categories))
		for k, cate := range categories {
			cates[k] = cate.Name
		}

		result[i] = RecipeListItemResponseDTO{
			ID:         entity.ID,
			Title:      entity.Name,
			PhotoURL:   entity.Image,
			Categories: cates,
		}
	}
	return &RecipeListResponseDTO{
		Items:      result,
		TotalItems: page.GetTotalItems(),
	}
}
