package mapper_impl

import (
  "p2/dto"
  "p2/models"
  "p2/service/mapper"
  "gorm.io/gorm"
)

type click struct {}

func NewClick() mapper.Click {
  return &click{}
}

func (m *click) ToDTO(entity models.Click) dto.ClickDTO {
  return dto.ClickDTO{
    ID:        entity.Model.ID,
    CreatedAt: entity.Model.CreatedAt,
    UpdatedAt: entity.Model.UpdatedAt,
    DeletedAt: entity.Model.DeletedAt,

    Ip: entity.Ip,
    Country: entity.Country,
    Referer: entity.Referer,
    RefererHost: entity.RefererHost,
    UserAgent: entity.UserAgent,

    LinkID: entity.LinkID,
  }
}

func (m *click) ToEntity(dto dto.ClickDTO) models.Click {
  return models.Click{
    Model: gorm.Model{
      ID:        dto.ID,
      CreatedAt: dto.CreatedAt,
      UpdatedAt: dto.UpdatedAt,
      DeletedAt: dto.DeletedAt,
    },

    Ip: dto.Ip,
    Country: dto.Country,
    Referer: dto.Referer,
    RefererHost: dto.RefererHost,
    UserAgent: dto.UserAgent,

    LinkID: dto.LinkID,
  }
}

func (m *click) ToDTOS(entityList []models.Click) []dto.ClickDTO {
  dtos := make([]dto.ClickDTO, len(entityList))

  for i, entity := range entityList {
    dtos[i] = m.ToDTO(entity)
  }

  return dtos
}

func (m *click) ToEntities(dtoList []dto.ClickDTO) []models.Click {
  entities := make([]models.Click, len(dtoList))

  for i, dto := range dtoList {
    entities[i] = m.ToEntity(dto)
  }

  return entities
}