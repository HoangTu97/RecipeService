package service

import (
	"Food/dto"
	"Food/repository"
	"Food/service/mapper"
)

func FindOneIngredientDTO(id uint) (dto.IngredientDTO, bool) {
	ingredient, err := repository.FindOneIngredient(id)
	if err != nil {
		return dto.IngredientDTO{}, false
	}
	return mapper.ToIngredientDTO(ingredient), true
}

func FindIngredientIDsByName(name string) []uint {
	ingredients := repository.FindIngredientByName(name)

	result := make([]uint, len(ingredients))
	for i, v := range ingredients {
		result[i] = v.ID
	}

	return result
}
