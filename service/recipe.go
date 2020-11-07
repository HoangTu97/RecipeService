package service

import (
	"Food/dto"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
	"Food/repository"
	"Food/service/mapper"
)

type Recipe interface {
	Save(recipeDTO dto.RecipeDTO) (dto.RecipeDTO, bool)
	FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page
	FindPageByCates(cates []models.Category, pageable pagination.Pageable) page.Page
	FindPageByName(name string, pageable pagination.Pageable) page.Page
	FindPageByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page
	FindPageByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page
	FindPage(pageable pagination.Pageable) page.Page
	FindIDsByName(name string) []uint
	FindOne(id uint) (dto.RecipeDTO, bool)
	FindOneWithCate(id uint) (models.Recipe, bool)
	CountByCateID(cateID uint) int64
}

type recipe struct {
	repository repository.Recipe
	mapper mapper.Recipe
}

func NewRecipe(repository repository.Recipe, mapper mapper.Recipe) Recipe {
	return &recipe{repository: repository, mapper: mapper}
}

func (s *recipe) Save(recipeDTO dto.RecipeDTO) (dto.RecipeDTO, bool) {
	recipe := s.mapper.ToEntity(recipeDTO)
	var err error
	recipe, err = s.repository.Save(recipe)
	if err != nil {
		return recipeDTO, false
	}
	return s.mapper.ToDTO(recipe), true
}

// FindPageByCateID return page models.Recipe
func (s *recipe) FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	return s.repository.FindPageByCateID(cateID, pageable)
}

func (s *recipe) FindPageByCates(cates []models.Category, pageable pagination.Pageable) page.Page {
	return s.repository.FindPageByCates(cates, pageable)
}

func (s *recipe) FindPageByName(name string, pageable pagination.Pageable) page.Page {
	return s.repository.FindPageByName(name, pageable)
}

func (s *recipe) FindPageByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page {
	return s.repository.FindPageByIngredientID(ingredientID, pageable)
}

func (s *recipe) FindPageByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page {
	return s.repository.FindPageByIngredientIDIn(ingredientIDs, pageable)
}

func (s *recipe) FindPage(pageable pagination.Pageable) page.Page {
	return s.repository.FindPage(pageable)
}

func (s *recipe) FindIDsByName(name string) []uint {
	recipes := s.repository.FindByName(name)
	ids := make([]uint, len(recipes))
	for i, v := range recipes {
		ids[i] = v.ID
	}
	return ids
}

func (s *recipe) FindOne(id uint) (dto.RecipeDTO, bool) {
	recipe, err := s.repository.FindOne(id)
	if err != nil {
		return dto.RecipeDTO{}, false
	}
	return s.mapper.ToDTO(recipe), true
}

func (s *recipe) FindOneWithCate(id uint) (models.Recipe, bool) {
	recipe, err := s.repository.FindOnePreloadCate(id)
	if err != nil {
		return models.Recipe{}, false
	}
	return recipe, true
}

func (s *recipe) CountByCateID(cateID uint) int64 {
	return s.repository.CountByCateID(cateID)
}
