package mapper

import (
  "p2/dto"
  "p2/models"
)

type Link interface {
  ToDTO(entity models.Link) dto.LinkDTO
  ToEntity(dto dto.LinkDTO) models.Link
  ToDTOS(entityList []models.Link) []dto.LinkDTO
  ToEntities(dtoList []dto.LinkDTO) []models.Link
}
