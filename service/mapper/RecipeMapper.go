package mapper

import (
	"Food/dto"
	"Food/models"
	"Food/pkg/converter"
)

type Recipe interface {
	ToDTO(entity models.Recipe) dto.RecipeDTO
	ToEntity(dto dto.RecipeDTO) models.Recipe
	ToDTOS(entityList []models.Recipe) []dto.RecipeDTO
	ToEntities(dtoList []dto.RecipeDTO) []models.Recipe
	ToDTOSInterfaceFromEntitiesInterface(interfaces []interface{}) []interface{}
}

type recipe struct {}

func NewRecipe() Recipe {
	return &recipe{}
}

func (m *recipe) ToDTO(entity models.Recipe) dto.RecipeDTO {
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

func (m *recipe) ToEntity(dto dto.RecipeDTO) models.Recipe {
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

func (m *recipe) ToDTOS(entityList []models.Recipe) []dto.RecipeDTO {
	dtos := make([]dto.RecipeDTO, len(entityList))

	for i, entity := range entityList {
		dtos[i] = m.ToDTO(entity)
	}

	return dtos
}

func (m *recipe) ToEntities(dtoList []dto.RecipeDTO) []models.Recipe {
	entities := make([]models.Recipe, len(dtoList))

	for i, dto := range dtoList {
		entities[i] = m.ToEntity(dto)
	}

	return entities
}

func (m *recipe) ToDTOSInterfaceFromEntitiesInterface(interfaces []interface{}) []interface{} {
	dtos := make([]interface{}, len(interfaces))

	for i, inter := range interfaces {
		dtos[i] = m.ToDTO(inter.(models.Recipe))
	}

	return dtos
}
