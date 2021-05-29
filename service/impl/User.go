package service_impl

import (
  "p2/pkg/domain"
  "p2/dto"
  "p2/helpers/page"
  "p2/helpers/pagination"
  "p2/pkg/service/Hash"
  "p2/repository"
  "p2/service"
  "p2/service/mapper"
)

type user struct {
  repository repository.User
  mapper     mapper.User
  hashService Hash.Service
}

func NewUser(repository repository.User, mapper mapper.User, hashService Hash.Service) service.User {
  return &user{repository: repository, mapper: mapper, hashService: hashService}
}

func (s *user) Create(userDTO dto.UserDTO) (dto.UserDTO, bool) {
  pass := s.hashService.Make(userDTO.Password)
  if pass == "" {
    return dto.UserDTO{}, false
  }
  userDTO.Password = pass
  userDTO.Roles = append(userDTO.Roles, domain.ROLE_USER)

  user := s.mapper.ToEntity(userDTO)

  var err error
  user, err = s.repository.Save(user)
  if err != nil {
    return userDTO, false
  }

  return s.mapper.ToDTO(user), true
}

func (s *user) Save(userDTO dto.UserDTO) (dto.UserDTO, bool) {
  user := s.mapper.ToEntity(userDTO)
  var err error
  user, err = s.repository.Save(user)
  if err != nil {
    return userDTO, false
  }
  return s.mapper.ToDTO(user), true
}

func (s *user) FindOneLogin(username string, password string) (dto.UserDTO, bool) {
  user, err := s.repository.FindOneByName(username)
  if err != nil {
    return dto.UserDTO{}, false
  }

  valid := s.hashService.Check(user.Password, password)
  if !valid {
    return dto.UserDTO{}, false
  }

  return s.mapper.ToDTO(user), true
}

func (s *user) FindOneByUserID(userId string) (dto.UserDTO, bool) {
  user, err := s.repository.FineOneByUserId(userId)
  if err != nil {
    return dto.UserDTO{}, false
  }

  return s.mapper.ToDTO(user), true
}

func (s *user) FindOneByUsername(username string) (dto.UserDTO, bool) {
  user, err := s.repository.FindOneByName(username)
  if err != nil {
    return dto.UserDTO{}, false
  }

  return s.mapper.ToDTO(user), true
}

func (s *user) FindPage(pageable pagination.Pageable) page.Page {
  return s.repository.FindPage(pageable)
}

func (s *user) Delete(id uint) {
  s.repository.Delete(id)
}
