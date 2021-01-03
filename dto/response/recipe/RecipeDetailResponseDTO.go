package recipe

import (
	"Food/models"
	"Food/pkg/converter"
)

// RecipeDetail_CateResponseDTO godoc
type RecipeDetail_CateResponseDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// RecipeDetailResponseDTO godoc
type RecipeDetailResponseDTO struct {
	ID          uint                           `json:"id"`
	Title       string                         `json:"title"`
	Photos      []string                       `json:"photos"`
	Duration    uint32                         `json:"duration"`
	Description string                         `json:"description"`
	Categories  []RecipeDetail_CateResponseDTO `json:"categories"`
	Steps       []string                       `json:"steps"`
}

// CreateRecipeDetailResponseDTO godoc
func CreateRecipeDetailResponseDTO(recipe models.Recipe) RecipeDetailResponseDTO {
	cates := make([]RecipeDetail_CateResponseDTO, len(recipe.Categories))

	for i, v := range recipe.Categories {
		cates[i] = RecipeDetail_CateResponseDTO{
			ID:   v.ID,
			Name: v.Name,
		}
	}

	return RecipeDetailResponseDTO{
		ID:          recipe.ID,
		Title:       recipe.Name,
		Photos:      converter.MustArrStr(recipe.Photos),
		Duration:    recipe.Duration,
		Description: recipe.Description,
		Categories:  cates,
		Steps:       converter.MustArrStr(recipe.Steps),
	}
}
