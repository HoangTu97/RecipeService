package repository_impl

import (
  "p2/helpers/page"
  "p2/helpers/pagination"
  "p2/models"
  "p2/repository"

  "gorm.io/gorm"
)

type click struct {
  db *gorm.DB
}

func NewClick(db *gorm.DB) repository.Click {
  return &click{db: db}
}

func (r *click) Save(click models.Click) (models.Click, error) {
  result := r.db.Save(&click)
  if result.Error != nil {
    return click, result.Error
  }
  return click, nil
}

func (r *click) FindOne(id uint) (models.Click, error) {
  var click models.Click

  result := r.db.First(&click, id)
  if result.Error != nil {
    return models.Click{}, result.Error
  }

  return click, nil
}

func (r *click) FindPage(pageable pagination.Pageable) page.Page {
  var clicks []models.Click

  paginator := pagination.Paging(&pagination.Param{
    DB:      r.db.Joins("User").Joins("Recipe"),
    Page:    pageable.GetPageNumber(),
    Limit:   pageable.GetPageSize(),
    ShowSQL: true,
  }, &clicks)

  return page.From(r.toInterfacesFromClick(clicks), paginator.TotalRecord)
}

func (r *click) toInterfacesFromClick(clicks []models.Click) []interface{} {
  content := make([]interface{}, len(clicks))
  for i, v := range clicks {
    content[i] = v
  }
  return content
}

func (r *click) Delete(id uint) {
  r.db.Delete(&models.Click{}, id)
}
