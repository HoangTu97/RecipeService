package repository

import (
	"Food/config"
	"Food/models"
)

func SaveUserRecipeInteraction(interaction models.UserRecipeInteraction) (models.UserRecipeInteraction, error) {
	result := config.GetDB().Save(&interaction)
	if result.Error != nil {
		return models.UserRecipeInteraction{}, result.Error
	}
	return interaction, nil
}

func FindOneUserRecipeInteraction(id uint) (models.UserRecipeInteraction, error) {
	var interaction models.UserRecipeInteraction

	result := config.GetDB().First(&interaction, id)
	if result.Error != nil {
		return models.UserRecipeInteraction{}, result.Error
	}

	// gredis.Set(key, category, 3600)
	return interaction, nil
}

func FindOneUserRecipeInteractionByUserIdAndRecipeId(userId uint, recipeId uint) (models.UserRecipeInteraction, error) {
	var interaction models.UserRecipeInteraction
	result := config.GetDB().Where("user_id = ? AND recipe_id = ?", userId, recipeId).First(&interaction)
	if result.Error != nil {
		return models.UserRecipeInteraction{}, result.Error
	}
	return interaction, nil
}

func FindAllUserRecipeInteraction() []models.UserRecipeInteraction {
	var interactions []models.UserRecipeInteraction
	config.GetDB().Find(&interactions)
	return interactions
}

func DeleteUserRecipeInteraction(interaction models.UserRecipeInteraction) error {
	result := config.GetDB().Delete(&interaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
