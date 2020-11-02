package repository

import (
	"Food/models"

	"gorm.io/gorm"
)

type user_recipe_interaction struct {
	db *gorm.DB
}

type UserRecipeInteraction interface {
	Save(interaction models.UserRecipeInteraction) (models.UserRecipeInteraction, error)
	FindOne(id uint) (models.UserRecipeInteraction, error)
	FindOneByUserIdAndRecipeId(userId uint, recipeId uint) (models.UserRecipeInteraction, error)
	FindAll() []models.UserRecipeInteraction
	Delete(interaction models.UserRecipeInteraction) error
}

func NewUserRecipeInteraction(db *gorm.DB) UserRecipeInteraction {
	return &user_recipe_interaction{db: db}
}

func (r *user_recipe_interaction) Save(interaction models.UserRecipeInteraction) (models.UserRecipeInteraction, error) {
	result := r.db.Save(&interaction)
	if result.Error != nil {
		return models.UserRecipeInteraction{}, result.Error
	}
	return interaction, nil
}

func (r *user_recipe_interaction) FindOne(id uint) (models.UserRecipeInteraction, error) {
	var interaction models.UserRecipeInteraction

	result := r.db.First(&interaction, id)
	if result.Error != nil {
		return models.UserRecipeInteraction{}, result.Error
	}

	// gredis.Set(key, category, 3600)
	return interaction, nil
}

func (r *user_recipe_interaction) FindOneByUserIdAndRecipeId(userId uint, recipeId uint) (models.UserRecipeInteraction, error) {
	var interaction models.UserRecipeInteraction
	result := r.db.Where("user_id = ? AND recipe_id = ?", userId, recipeId).First(&interaction)
	if result.Error != nil {
		return models.UserRecipeInteraction{}, result.Error
	}
	return interaction, nil
}

func (r *user_recipe_interaction) FindAll() []models.UserRecipeInteraction {
	var interactions []models.UserRecipeInteraction
	r.db.Find(&interactions)
	return interactions
}

func (r *user_recipe_interaction) Delete(interaction models.UserRecipeInteraction) error {
	result := r.db.Delete(&interaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
