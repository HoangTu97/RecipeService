package service

import (
	"Food/dto"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
)

type categoryProxy struct {
	service category
}

func NewCategoryProxy(service category) Category {
	return &categoryProxy{service: service}
}

func (s *categoryProxy) Save(categoryDTO dto.CategoryDTO) (dto.CategoryDTO, bool) {
	return s.service.Save(categoryDTO)
}

func (s *categoryProxy) FindOne(id uint) (dto.CategoryDTO, bool) {
	return s.service.FindOne(id)
}

func (s *categoryProxy) FindByName(name string) ([]models.Category, bool) {
	return s.service.FindByName(name)
}

func (s *categoryProxy) FindAll() []dto.CategoryDTO {
	return s.service.FindAll()
}

func (s *categoryProxy) FindPage(pageable pagination.Pageable) page.Page {
	return s.service.FindPage(pageable)
}

func (s *categoryProxy) Delete(id uint) bool {
	return s.service.Delete(id)
}
