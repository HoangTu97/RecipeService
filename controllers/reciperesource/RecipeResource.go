package reciperesource

import (
	"Food/dto/response"
	"Food/dto/response/recipe"
	"Food/helpers/converter"
	"Food/helpers/pagination"
	"Food/service"

	"github.com/gin-gonic/gin"
)

// Recipe GetByCategory
// @Summary GetByCategory
// @Tags PublicRecipe
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Param categoryId path int true "categoryId"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeListResponseDTO} "desc"
// @Router /api/public/recipe/getByCategory/:categoryId [get]
func GetByCategory(c *gin.Context) {
	pageable := pagination.GetPage(c)
	id := converter.MustUint(c.Param("categoryId"))

	_, isExist := service.FindOneCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	page := service.FindPageRecipeByCateID(id, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

// Recipe GetCountByCategory
// @Summary GetCountByCategory
// @Tags PublicRecipe
// @Accept json
// @Param categoryId path int true "categoryId"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeCountByCateResponseDTO} "desc"
// @Router /api/public/recipe/countByCategory/:categoryId [get]
func GetCountByCategory(c *gin.Context) {
	id := converter.MustUint(c.Param("categoryId"))

	_, isExist := service.FindOneCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	result := service.CountRecipeByCateID(id)
	response.CreateSuccesResponse(c, &recipe.RecipeCountByCateResponseDTO{
		Value: result,
	})
}

func GetByCategoryName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	cateName := converter.MustString(c.Query("name"))

	cates, isExist := service.FindCateByName(cateName)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	page := service.FindPageRecipeByCates(cates, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func GetByIngredient(c *gin.Context) {
	pageable := pagination.GetPage(c)
	id := converter.MustUint(c.Param("ingredientId"))

	_, isExist := service.FindOneIngredientDTO(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}

	page := service.FindPageRecipeByIngredientID(id, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

func GetByIngredientName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	ingredientName := converter.MustString(c.Query("name"))

	ingredientIDs := service.FindIngredientIDsByName(ingredientName)
	if len(ingredientIDs) == 0 {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}

	page := service.FindPageRecipeByIngredientIDIn(ingredientIDs, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

// Recipe getByRecipeName
// @Summary GetByRecipeName
// @Tags PublicRecipe
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Param name query string true "name"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeListResponseDTO} "desc"
// @Router /api/public/recipe/searchByRecipeName [get]
func GetByRecipeName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	recipeName := converter.MustString(c.Query("name"))

	page := service.FindPageRecipeByName(recipeName, pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

// Recipe all
// @Summary GetAll
// @Tags PublicRecipe
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeListResponseDTO} "desc"
// @Router /api/public/recipe/getAll [get]
func GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := service.FindPageRecipe(pageable)

	response.CreateSuccesResponse(c, recipe.CreateRecipeListResponseDTOFromPage(page))
}

// Recipe getDetailByID
// @Summary GetDetailByID
// @Tags PublicRecipe
// @Accept json
// @Param id path int true "id"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeDetailResponseDTO} "desc"
// @Router /api/public/recipe/detail/:id [get]
func GetDetailByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))

	recipeData, isExist := service.FindOneRecipeWithCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "RECIPE_NOT_FOUND")
		return
	}

	response.CreateSuccesResponse(c, recipe.CreateRecipeDetailResponseDTO(recipeData))
}
