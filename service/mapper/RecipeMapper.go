package mapper

import (
	"Food/dto"
	"Food/helpers/converter"
	"Food/models"
)

func ToRecipeDTO(entity models.Recipe) dto.RecipeDTO {
	return dto.RecipeDTO{
		ID:          entity.ID,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
		DeletedAt:   entity.DeletedAt,
		Name:        entity.Name,
		Image:       entity.Image,
		Description: entity.Description,
		Duration:    entity.Duration,
		Photos:      converter.MustArrStr(entity.Photos),
	}
}

func ToRecipe(dto dto.RecipeDTO) models.Recipe {
	return models.Recipe{
		ID:          dto.ID,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
		DeletedAt:   dto.DeletedAt,
		Name:        dto.Name,
		Image:       dto.Image,
		Description: dto.Description,
		Duration:    dto.Duration,
		Photos:      converter.ToStr(dto.Photos),
	}
}

func ToRecipeDTOS(entityList []models.Recipe) []dto.RecipeDTO {
	dtos := make([]dto.RecipeDTO, len(entityList))

	for i, entity := range entityList {
		dtos[i] = ToRecipeDTO(entity)
	}

	return dtos
}

func ToRecipes(dtoList []dto.RecipeDTO) []models.Recipe {
	entities := make([]models.Recipe, len(dtoList))

	for i, dto := range dtoList {
		entities[i] = ToRecipe(dto)
	}

	return entities
}

func ToDTOSInterfaceFromEntitiesInterface(interfaces []interface{}) []interface{} {
	dtos := make([]interface{}, len(interfaces))

	for i, inter := range interfaces {
		dtos[i] = ToRecipeDTO(inter.(models.Recipe))
	}

	return dtos
}
