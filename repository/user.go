package repository

import (
	"Food/models"

	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

type User interface {
	Save(user models.User) models.User
	FineOneByUserId(userId string) (models.User, error)
	FindOneByName(name string) (models.User, error)
}

func NewUser(db *gorm.DB) User {
	return &user{ db: db}
}

func (r *user) Save(user models.User) models.User {
	r.db.Create(&user)
	return user
}

func (r *user) FineOneByUserId(userId string) (models.User, error) {
	user := models.User{}
	result := r.db.Where("user_id = ?", userId).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *user) FindOneByName(name string) (models.User, error) {
	user := models.User{}
	result := r.db.Where("name = ?", name).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}