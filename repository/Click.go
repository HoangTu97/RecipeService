package repository

import (
  "p2/helpers/page"
  "p2/helpers/pagination"
  "p2/models"
)

type Click interface {
  Save(click models.Click) (models.Click, error)
  FindOne(id uint) (models.Click, error)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
}
