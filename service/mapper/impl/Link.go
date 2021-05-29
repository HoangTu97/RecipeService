package mapper_impl

import (
  "p2/dto"
  "p2/models"
  "p2/service/mapper"
  "gorm.io/gorm"
)

type link struct {}

func NewLink() mapper.Link {
  return &link{}
}

func (m *link) ToDTO(entity models.Link) dto.LinkDTO {
  return dto.LinkDTO{
    ID:        entity.Model.ID,
    CreatedAt: entity.Model.CreatedAt,
    UpdatedAt: entity.Model.UpdatedAt,
    DeletedAt: entity.Model.DeletedAt,

    ShortUrl: entity.ShortUrl,
    LongUrl: entity.LongUrl,
    LongUrlHash: entity.LongUrlHash,
    Ip: entity.Ip,
    ClickNum: entity.ClickNum,
    SecretKey: entity.SecretKey,
    IsDisabled: entity.IsDisabled,
    IsCustom: entity.IsCustom,
    IsApi: entity.IsApi,

    CreatorID: entity.CreatorID,
  }
}

func (m *link) ToEntity(dto dto.LinkDTO) models.Link {
  return models.Link{
    Model: gorm.Model{
      ID:        dto.ID,
      CreatedAt: dto.CreatedAt,
      UpdatedAt: dto.UpdatedAt,
      DeletedAt: dto.DeletedAt,
    },

    ShortUrl: dto.ShortUrl,
    LongUrl: dto.LongUrl,
    LongUrlHash: dto.LongUrlHash,
    Ip: dto.Ip,
    ClickNum: dto.ClickNum,
    SecretKey: dto.SecretKey,
    IsDisabled: dto.IsDisabled,
    IsCustom: dto.IsCustom,
    IsApi: dto.IsApi,

    CreatorID: dto.CreatorID,
  }
}

func (m *link) ToDTOS(entityList []models.Link) []dto.LinkDTO {
  dtos := make([]dto.LinkDTO, len(entityList))

  for i, entity := range entityList {
    dtos[i] = m.ToDTO(entity)
  }

  return dtos
}

func (m *link) ToEntities(dtoList []dto.LinkDTO) []models.Link {
  entities := make([]models.Link, len(dtoList))

  for i, dto := range dtoList {
    entities[i] = m.ToEntity(dto)
  }

  return entities
}