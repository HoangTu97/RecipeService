package service

import (
	"Food/dto"
	"Food/helpers/cache"
	"encoding/json"
)

type userProxy struct {
	service User
	cache cache.Cache
}

func NewUserProxy(service User, cache cache.Cache) User {
	return &userProxy{service: service, cache: cache}
}

func (s *userProxy) Create(userDTO dto.UserDTO) (dto.UserDTO, bool) {
	return s.service.Create(userDTO)
}

func (s *userProxy) FindOneLogin(username string, password string) (dto.UserDTO, bool) {
	var userDTO dto.UserDTO

	key := s.cache.GenKey("UserDTO", username, password)
	if s.cache.Exists(key) {
		data, err := s.cache.Get(key)
		if err == nil {
			return dto.UserDTO{}, false
		}
		_ = json.Unmarshal(data, &userDTO)
		return userDTO, true
	}

	userDTO, exist := s.service.FindOneLogin(username, password)
	if !exist {
		return dto.UserDTO{}, false
	}

	_ = s.cache.Set(key, userDTO, 3600)

	return userDTO, true
}

func (s *userProxy) FindOneByUserID(userId string) (dto.UserDTO, bool) {
	return s.service.FindOneByUserID(userId)
}
