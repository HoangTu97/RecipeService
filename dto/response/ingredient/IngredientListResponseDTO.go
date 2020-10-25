package ingredient

import (
	"Food/models"
)

// IngredientListResponseDTO godoc
type IngredientListResponseDTO struct {
	Items      []IngredientListItemResponseDTO `json:"items"`
	TotalItems int                             `json:"totalItems"`
}

// CreateEmptyIngredientListResponseDTO godoc
func CreateEmptyIngredientListResponseDTO() *IngredientListResponseDTO {
	return &IngredientListResponseDTO{
		Items:      []IngredientListItemResponseDTO{},
		TotalItems: 0,
	}
}

// CreateIngredientListResponseDTO create list response from list models.RecipeIngredients
func CreateIngredientListResponseDTO(data []models.RecipeIngredients) *IngredientListResponseDTO {
	result := make([]IngredientListItemResponseDTO, len(data))
	for i, v := range data {
		result[i] = IngredientListItemResponseDTO{
			ID:       v.Ingredient.ID,
			Name:     v.Ingredient.Name,
			Image:    v.Ingredient.Image,
			Quantity: v.Quantity,
		}
	}
	return &IngredientListResponseDTO{
		Items:      result,
		TotalItems: len(result),
	}
}
