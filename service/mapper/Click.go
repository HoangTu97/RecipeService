package mapper

import (
  "p2/dto"
  "p2/models"
)

type Click interface {
  ToDTO(entity models.Click) dto.ClickDTO
  ToEntity(dto dto.ClickDTO) models.Click
  ToDTOS(entityList []models.Click) []dto.ClickDTO
  ToEntities(dtoList []dto.ClickDTO) []models.Click
}
