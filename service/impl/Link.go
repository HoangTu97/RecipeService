package service_impl

import (
  "p2/dto"
  "p2/helpers/page"
  "p2/helpers/pagination"
  "p2/repository"
  "p2/service"
  "p2/service/mapper"
  "math"
)

type link struct {
  repository repository.Link
  mapper mapper.Link
}

func NewLink(repository repository.Link, mapper mapper.Link) service.Link {
  return &link{repository: repository, mapper: mapper}
}

func (s *link) Save(linkDTO dto.LinkDTO) (dto.LinkDTO, bool) {
  link := s.mapper.ToEntity(linkDTO)
  var err error
  link, err = s.repository.Save(link)
  if err != nil {
    return linkDTO, false
  }
  return s.mapper.ToDTO(link), true
}

func (s *link) FindOne(id uint) (dto.LinkDTO, bool) {
  link, err := s.repository.FindOne(id)
  if err != nil {
    return dto.LinkDTO{}, false
  }
  return s.mapper.ToDTO(link), true
}

func (s *link) FindPage(pageable pagination.Pageable) page.Page {
  return s.repository.FindPage(pageable)
}

func (s *link) Delete(id uint) {
  s.repository.Delete(id)
}

var mapChar = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (s *link) IdToShortUrl(id uint) string {
  short_url := make([]rune, 50)
  n := 0
  for id > 0 {
    short_url[n] = mapChar[int(math.Mod(float64(id), 62))]
    id = id / 62
    n++
  }
  short_url = short_url[0:n]
  // Reverse 
  for i := 0; i < n/2; i++ { 
    short_url[i], short_url[n-1-i] = short_url[n-1-i], short_url[i]
  }
  return string(short_url)
}

func (s *link) ShortUrlToId(short_url string) uint {
  var id uint = 0
  // Base decode conversion logic.
  for i := 0; i < len(short_url); i++ {
    if 'a' <= short_url[i] && short_url[i] <= 'z' {
      id = id * 62 + (uint(short_url[i]) - 97)
    }
    if ('A' <= short_url[i] && short_url[i] <= 'Z') {
      id = id * 62 + (uint(short_url[i]) - 65) + 26
    }
    if ('0' <= short_url[i] && short_url[i] <= '9') {
      id = id * 62 + (uint(short_url[i]) - 48) + 52
    }
  }
  return id
}
