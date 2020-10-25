package service

import (
	"Food/dto"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
	"Food/repository"
	"Food/service/mapper"
)

func SaveRecipe(recipeDTO dto.RecipeDTO) (dto.RecipeDTO, bool) {
	recipe := mapper.ToRecipe(recipeDTO)
	var err error
	recipe, err = repository.SaveRecipe(recipe)
	if err != nil {
		return recipeDTO, false
	}
	return mapper.ToRecipeDTO(recipe), true
}

// FindPageByCateID return page models.Recipe
func FindPageRecipeByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	return repository.FindPageRecipeByCateID(cateID, pageable)
}

func FindPageRecipeByCates(cates []models.Category, pageable pagination.Pageable) page.Page {
	return repository.FindPageRecipeByCates(cates, pageable)
}

func FindPageRecipeByName(name string, pageable pagination.Pageable) page.Page {
	return repository.FindPageRecipeByName(name, pageable)
}

func FindPageRecipeByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page {
	return repository.FindPageRecipeByIngredientID(ingredientID, pageable)
}

func FindPageRecipeByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page {
	return repository.FindPageRecipeByIngredientIDIn(ingredientIDs, pageable)
}

func FindPageRecipe(pageable pagination.Pageable) page.Page {
	return repository.FindPageRecipe(pageable)
}

func FindRecipeIDsByName(name string) []uint {
	recipes := repository.FindRecipeByName(name)
	ids := make([]uint, len(recipes))
	for i, v := range recipes {
		ids[i] = v.ID
	}
	return ids
}

func FindOneRecipe(id uint) (dto.RecipeDTO, bool) {
	recipe, err := repository.FindOneRecipe(id)
	if err != nil {
		return dto.RecipeDTO{}, false
	}
	return mapper.ToRecipeDTO(recipe), true
}

func FindOneRecipeWithCate(id uint) (models.Recipe, bool) {
	recipe, err := repository.FindOneRecipePreloadCate(id)
	if err != nil {
		return models.Recipe{}, false
	}
	return recipe, true
}

func CountRecipeByCateID(cateID uint) int64 {
	return repository.CountRecipeByCateID(cateID)
}
