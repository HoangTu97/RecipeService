package service

import (
  "p2/dto"
  "p2/helpers/page"
  "p2/helpers/pagination"
)

type Link interface {
  Save(linkDTO dto.LinkDTO) (dto.LinkDTO, bool)
  FindOne(id uint) (dto.LinkDTO, bool)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
  IdToShortUrl(id uint) string
  ShortUrlToId(url string) uint
}
