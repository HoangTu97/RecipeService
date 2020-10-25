package repository

import (
	"Food/config"
	"Food/models"
)

func SaveUser(user models.User) models.User {
	config.GetDB().Create(&user)
	return user
}

func FineOneUserByUserId(userId string) (models.User, error) {
	user := models.User{}
	result := config.GetDB().Where("user_id = ?", userId).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func FindOneUserByName(name string) (models.User, error) {
	user := models.User{}
	result := config.GetDB().Where("name = ?", name).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}