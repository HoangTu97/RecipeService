package service

import (
	"Food/domain"
	"Food/dto"
	"Food/helpers/logging"
	"Food/repository"
	"Food/service/mapper"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(userDTO dto.UserDTO) (dto.UserDTO, bool) {
	pass, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		logging.Error(err)
		return dto.UserDTO{}, false
	}
	userDTO.Password = string(pass)
	userDTO.Roles = append(userDTO.Roles, domain.ROLE_USER)

	user := mapper.ToUser(userDTO)
	repository.SaveUser(user)

	return mapper.ToUserDTO(user), true
}

func FindOneUserLogin(username string, password string) (dto.UserDTO, bool) {
	user, err := repository.FindOneUserByName(username)
	if err != nil {
		return dto.UserDTO{}, false
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return dto.UserDTO{}, false
	}

	return mapper.ToUserDTO(user), true
}

func FindOneUserByUserID(userId string) (dto.UserDTO, bool) {
	user, err := repository.FineOneUserByUserId(userId)
	if err != nil {
		return dto.UserDTO{}, false
	}

	return mapper.ToUserDTO(user), true
}
