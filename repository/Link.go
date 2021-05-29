package repository

import (
  "p2/helpers/page"
  "p2/helpers/pagination"
  "p2/models"
)

type Link interface {
  Save(link models.Link) (models.Link, error)
  FindOne(id uint) (models.Link, error)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
}
