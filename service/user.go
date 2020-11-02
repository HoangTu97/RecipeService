package service

import (
	"Food/domain"
	"Food/dto"
	"Food/helpers/logging"
	"Food/repository"
	"Food/service/mapper"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Create(userDTO dto.UserDTO) (dto.UserDTO, bool)
	FindOneLogin(username string, password string) (dto.UserDTO, bool)
	FindOneByUserID(userId string) (dto.UserDTO, bool)
}

type user struct {
	repository repository.User
}

func NewUser(repository repository.User) User {
	return &user{repository: repository}
}

func (s *user) Create(userDTO dto.UserDTO) (dto.UserDTO, bool) {
	pass, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		logging.Error(err)
		return dto.UserDTO{}, false
	}
	userDTO.Password = string(pass)
	userDTO.Roles = append(userDTO.Roles, domain.ROLE_USER)

	user := mapper.ToUser(userDTO)
	s.repository.Save(user)

	return mapper.ToUserDTO(user), true
}

func (s *user) FindOneLogin(username string, password string) (dto.UserDTO, bool) {
	user, err := s.repository.FindOneByName(username)
	if err != nil {
		return dto.UserDTO{}, false
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return dto.UserDTO{}, false
	}

	return mapper.ToUserDTO(user), true
}

func (s *user) FindOneByUserID(userId string) (dto.UserDTO, bool) {
	user, err := s.repository.FineOneByUserId(userId)
	if err != nil {
		return dto.UserDTO{}, false
	}

	return mapper.ToUserDTO(user), true
}
