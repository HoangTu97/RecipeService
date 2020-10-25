package ingredientresource

import (
	"Food/dto/response"
	"Food/dto/response/ingredient"
	"Food/helpers/converter"
	"Food/service"

	"github.com/gin-gonic/gin"
)

func GetNameByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	ingredientDTO, isExist := service.FindOneIngredientDTO(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}
	response.CreateSuccesResponse(c, &ingredient.IngredientNameResponseDTO{Name: ingredientDTO.Name})
}

func GetImageByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	ingredientDTO, isExist := service.FindOneIngredientDTO(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}
	response.CreateSuccesResponse(c, &ingredient.IngredientImageResponseDTO{Image: ingredientDTO.Image})
}

func GetByRecipeName(c *gin.Context) {
	recipeName := converter.MustString(c.Query("recipeName"))
	recipeIDs := service.FindRecipeIDsByName(recipeName)

	if len(recipeIDs) == 0 {
		response.CreateSuccesResponse(c, ingredient.CreateEmptyIngredientListResponseDTO())
		return
	}

	recipeIngredients := service.FindRecipeIngredientsByRecipeIDs(recipeIDs)
	response.CreateSuccesResponse(c, ingredient.CreateIngredientListResponseDTO(recipeIngredients))
}

func GetByRecipeID(c *gin.Context) {
	recipeID := converter.MustUint(c.Query("recipeId"))
	_, isExist := service.FindOneRecipe(recipeID)
	if !isExist {
		response.CreateErrorResponse(c, "RECIPE_NOT_FOUND")
		return
	}

	recipeIngredients := service.FindRecipeIngredientsByRecipeID(recipeID)
	response.CreateSuccesResponse(c, ingredient.CreateIngredientListResponseDTO(recipeIngredients))
}
