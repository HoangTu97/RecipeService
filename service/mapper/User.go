package mapper

import (
  "p2/dto"
  "p2/models"
)

type User interface {
  ToDTO(entity models.User) dto.UserDTO
  ToEntity(dto dto.UserDTO) models.User
  ToDTOS(entityList []models.User) []dto.UserDTO
  ToEntities(dtoList []dto.UserDTO) []models.User
}
