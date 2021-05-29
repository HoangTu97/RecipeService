package service_impl

import (
  "p2/dto"
  "p2/helpers/page"
  "p2/helpers/pagination"
  "p2/repository"
  "p2/service"
  "p2/service/mapper"
)

type click struct {
  repository repository.Click
  mapper mapper.Click
}

func NewClick(repository repository.Click, mapper mapper.Click) service.Click {
  return &click{repository: repository, mapper: mapper}
}

func (s *click) Save(clickDTO dto.ClickDTO) (dto.ClickDTO, bool) {
  click := s.mapper.ToEntity(clickDTO)
  var err error
  click, err = s.repository.Save(click)
  if err != nil {
    return clickDTO, false
  }
  return s.mapper.ToDTO(click), true
}

func (s *click) FindOne(id uint) (dto.ClickDTO, bool) {
  click, err := s.repository.FindOne(id)
  if err != nil {
    return dto.ClickDTO{}, false
  }
  return s.mapper.ToDTO(click), true
}

func (s *click) FindPage(pageable pagination.Pageable) page.Page {
  return s.repository.FindPage(pageable)
}

func (s *click) Delete(id uint) {
  s.repository.Delete(id)
}
