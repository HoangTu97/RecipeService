package service

import (
	"Food/domain"
	"Food/dto"
	"Food/helpers/jwt"
	"Food/repository"
	"Food/service/mapper"

	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Create(userDTO dto.UserDTO) (dto.UserDTO, bool)
	GetUserToken(userDTO dto.UserDTO) (string, error)
	FindOneLogin(username string, password string) (dto.UserDTO, bool)
	FindOneByUserID(userId string) (dto.UserDTO, bool)
	FindOneByUsername(username string) (dto.UserDTO, bool)
}

type user struct {
	repository repository.User
	mapper     mapper.User
}

func NewUser(repository repository.User, mapper mapper.User) User {
	return &user{repository: repository, mapper: mapper}
}

func (s *user) Create(userDTO dto.UserDTO) (dto.UserDTO, bool) {
	pass, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserDTO{}, false
	}
	userDTO.Password = string(pass)
	userDTO.Roles = append(userDTO.Roles, domain.ROLE_USER)

	user := s.mapper.ToEntity(userDTO)
	s.repository.Save(user)

	return s.mapper.ToDTO(user), true
}

func (s *user) GetUserToken(userDTO dto.UserDTO) (string, error) {
	tokenString, err := jwt.GenerateToken(userDTO.UserID, userDTO.Name, userDTO.GetRolesStr())
	if err != nil {
		return "", err
	}

	return tokenString, nil
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
