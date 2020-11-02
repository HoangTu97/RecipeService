package service

import (
	"Food/dto"
	"Food/repository"
	"Food/service/mapper"
)

type Ingredient interface {
	FindOneDTO(id uint) (dto.IngredientDTO, bool)
	FindIDsByName(name string) []uint
}

type ingredient struct {
	repository repository.Ingredient
}

func NewIngredient(repository repository.Ingredient) Ingredient {
	return &ingredient{repository: repository}
}

func (s *ingredient) FindOneDTO(id uint) (dto.IngredientDTO, bool) {
	ingredient, err := s.repository.FindOne(id)
	if err != nil {
		return dto.IngredientDTO{}, false
	}
	return mapper.ToIngredientDTO(ingredient), true
}

func (s *ingredient) FindIDsByName(name string) []uint {
	ingredients := s.repository.FindByName(name)

	result := make([]uint, len(ingredients))
	for i, v := range ingredients {
		result[i] = v.ID
	}

	return result
}
