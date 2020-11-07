package service

import (
	"Food/dto"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
	"Food/repository"
	"Food/service/mapper"
)

type Category interface {
	Save(categoryDTO dto.CategoryDTO) (dto.CategoryDTO, bool)
	FindOne(id uint) (dto.CategoryDTO, bool)
	FindByName(name string) ([]models.Category, bool)
	FindAll() []dto.CategoryDTO
	FindPage(pageable pagination.Pageable) page.Page
	Delete(id uint) bool
}

type category struct {
	repository repository.Category
	mapper mapper.Category
}

func NewCategory(repository repository.Category, mapper mapper.Category) Category {
	return &category{repository: repository, mapper: mapper}
}

func (s *category) Save(categoryDTO dto.CategoryDTO) (dto.CategoryDTO, bool) {
	category := s.mapper.ToEntity(categoryDTO)
	var err error
	category, err = s.repository.Save(category)
	if err != nil {
		return categoryDTO, false
	}
	return s.mapper.ToDTO(category), true
}

func (s *category) FindOne(id uint) (dto.CategoryDTO, bool) {
	category, err := s.repository.FindOne(id)
	if err != nil {
		return dto.CategoryDTO{}, false
	}
	return s.mapper.ToDTO(category), true
}

func (s *category) FindByName(name string) ([]models.Category, bool) {
	categories, err := s.repository.FindOneByName(name)
	if err != nil {
		return []models.Category{}, false
	}
	return categories, true
}

func (s *category) FindAll() []dto.CategoryDTO {
	categories := s.repository.FindAll()
	return s.mapper.ToDTOS(categories)
}

func (s *category) FindPage(pageable pagination.Pageable) page.Page {
	return s.repository.FindPage(pageable)
}

func (s *category) Delete(id uint) bool {
	category, err := s.repository.FindOne(id)
	if err != nil {
		return false
	}
	_ = s.repository.Delete(category)
	return true
}
