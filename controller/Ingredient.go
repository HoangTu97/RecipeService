package controller

import (
	"Food/dto/response"
	IngreResponse "Food/dto/response/ingredient"
	"Food/helpers/converter"
	"Food/service"

	"github.com/gin-gonic/gin"
)

type Ingredient interface {
	GetNameByID(c *gin.Context)
	GetImageByID(c *gin.Context)
	GetByRecipeName(c *gin.Context)
	GetByRecipeID(c *gin.Context)
}

type ingredient struct {
	service service.Ingredient
	recipeService service.Recipe
	recipeIngredientService service.RecipeIngredients
}

func NewIngredient(service service.Ingredient, recipeService service.Recipe, recipeIngredientService service.RecipeIngredients) Ingredient {
	return &ingredient{service: service, recipeService: recipeService, recipeIngredientService: recipeIngredientService}
}

func (r *ingredient) GetNameByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	ingredientDTO, isExist := r.service.FindOneDTO(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}
	response.CreateSuccesResponse(c, &IngreResponse.IngredientNameResponseDTO{Name: ingredientDTO.Name})
}

func (r *ingredient) GetImageByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))
	ingredientDTO, isExist := r.service.FindOneDTO(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}
	response.CreateSuccesResponse(c, &IngreResponse.IngredientImageResponseDTO{Image: ingredientDTO.Image})
}

func (r *ingredient) GetByRecipeName(c *gin.Context) {
	recipeName := converter.MustString(c.Query("recipeName"))
	recipeIDs := r.recipeService.FindIDsByName(recipeName)

	if len(recipeIDs) == 0 {
		response.CreateSuccesResponse(c, IngreResponse.CreateEmptyIngredientListResponseDTO())
		return
	}

	recipeIngredients := r.recipeIngredientService.FindByRecipeIDs(recipeIDs)
	response.CreateSuccesResponse(c, IngreResponse.CreateIngredientListResponseDTO(recipeIngredients))
}

func (r *ingredient) GetByRecipeID(c *gin.Context) {
	recipeID := converter.MustUint(c.Query("recipeId"))
	_, isExist := r.recipeService.FindOne(recipeID)
	if !isExist {
		response.CreateErrorResponse(c, "RECIPE_NOT_FOUND")
		return
	}

	recipeIngredients := r.recipeIngredientService.FindByRecipeID(recipeID)
	response.CreateSuccesResponse(c, IngreResponse.CreateIngredientListResponseDTO(recipeIngredients))
}
