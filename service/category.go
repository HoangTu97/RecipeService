package service

import (
	"Food/dto"
	"Food/models"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/repository"
	"Food/service/mapper"
)

func SaveCate(categoryDTO dto.CategoryDTO) (dto.CategoryDTO, bool) {
	category := mapper.ToCategory(categoryDTO)
	var err error
	category, err = repository.SaveCate(category)
	if err != nil {
		return categoryDTO, false
	}
	return mapper.ToCategoryDTO(category), true
}

func FindOneCate(id uint) (dto.CategoryDTO, bool) {
	category, err := repository.FindOneCate(id)
	if err != nil {
		return dto.CategoryDTO{}, false
	}
	return mapper.ToCategoryDTO(category), true
}

func FindCateByName(name string) ([]models.Category, bool) {
	categories, err := repository.FindOneCateByName(name)
	if err != nil {
		return []models.Category{}, false
	}
	return categories, true
}

func FindAllCate() []dto.CategoryDTO {
	categories := repository.FindAllCate()
	return mapper.ToCategoryDTOS(categories)
}

func FindPageCate(pageable pagination.Pageable) page.Page {
	return repository.FindPageCate(pageable)
}

func DeleteCate(id uint) bool {
	category, err := repository.FindOneCate(id)
	if err != nil {
		return false
	}
	_ = repository.DeleteCate(category)
	return true
}
