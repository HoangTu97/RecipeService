package repository_impl

import (
  "p2/helpers/page"
  "p2/helpers/pagination"
  "p2/models"
  "p2/repository"

  "gorm.io/gorm"
)

type link struct {
  db *gorm.DB
}

func NewLink(db *gorm.DB) repository.Link {
  return &link{db: db}
}

func (r *link) Save(link models.Link) (models.Link, error) {
  result := r.db.Save(&link)
  if result.Error != nil {
    return link, result.Error
  }
  return link, nil
}

func (r *link) FindOne(id uint) (models.Link, error) {
  var link models.Link

  result := r.db.First(&link, id)
  if result.Error != nil {
    return models.Link{}, result.Error
  }

  return link, nil
}

func (r *link) FindPage(pageable pagination.Pageable) page.Page {
  var links []models.Link

  paginator := pagination.Paging(&pagination.Param{
    DB:      r.db.Joins("User").Joins("Recipe"),
    Page:    pageable.GetPageNumber(),
    Limit:   pageable.GetPageSize(),
    ShowSQL: true,
  }, &links)

  return page.From(r.toInterfacesFromLink(links), paginator.TotalRecord)
}

func (r *link) toInterfacesFromLink(links []models.Link) []interface{} {
  content := make([]interface{}, len(links))
  for i, v := range links {
    content[i] = v
  }
  return content
}

func (r *link) Delete(id uint) {
  r.db.Delete(&models.Link{}, id)
}
