package service

import (
  "p2/dto"
  "p2/helpers/page"
  "p2/helpers/pagination"
)

type Click interface {
  Save(clickDTO dto.ClickDTO) (dto.ClickDTO, bool)
  FindOne(id uint) (dto.ClickDTO, bool)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
}
