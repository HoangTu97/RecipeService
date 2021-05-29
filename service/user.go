package service

import (
  "p2/dto"
  "p2/helpers/page"
  "p2/helpers/pagination"
)

type User interface {
  Create(userDTO dto.UserDTO) (dto.UserDTO, bool)
  Save(userDTO dto.UserDTO) (dto.UserDTO, bool)
  FindOneLogin(username string, password string) (dto.UserDTO, bool)
  FindOneByUserID(userId string) (dto.UserDTO, bool)
  FindOneByUsername(username string) (dto.UserDTO, bool)
  FindPage(pageable pagination.Pageable) page.Page
  Delete(id uint)
}
