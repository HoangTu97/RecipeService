package controller

import (
	"Food/dto/response"
	RecipeResponse "Food/dto/response/recipe"
	"Food/helpers/pagination"
	"Food/pkg/converter"
	"Food/service"

	"github.com/gin-gonic/gin"
)

type Recipe interface {
	GetByCategory(c *gin.Context)
	GetCountByCategory(c *gin.Context)
	GetByCategoryName(c *gin.Context)
	GetByIngredient(c *gin.Context)
	GetByIngredientName(c *gin.Context)
	GetByRecipeName(c *gin.Context)
	GetAll(c *gin.Context)
	GetDetailByID(c *gin.Context)
}

type recipe struct {
	service service.Recipe
	cateService service.Category
	ingreService service.Ingredient
}

func NewRecipe(service service.Recipe, cateService service.Category, ingreService service.Ingredient) Recipe {
	return &recipe{
		service: service, 
		cateService: cateService,
		ingreService: ingreService,
	}
}

// Recipe GetByCategory
// @Summary GetByCategory
// @Tags PublicRecipe
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Param categoryId path int true "categoryId"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeListResponseDTO} "desc"
// @Router /api/public/recipe/getByCategory/:categoryId [get]
func (r *recipe) GetByCategory(c *gin.Context) {
	pageable := pagination.GetPage(c)
	id := converter.MustUint(c.Param("categoryId"))

	_, isExist := r.cateService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	page := r.service.FindPageByCateID(id, pageable)

	response.CreateSuccesResponse(c, RecipeResponse.CreateRecipeListResponseDTOFromPage(page))
}

// Recipe GetCountByCategory
// @Summary GetCountByCategory
// @Tags PublicRecipe
// @Accept json
// @Param categoryId path int true "categoryId"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeCountByCateResponseDTO} "desc"
// @Router /api/public/recipe/countByCategory/:categoryId [get]
func (r *recipe) GetCountByCategory(c *gin.Context) {
	id := converter.MustUint(c.Param("categoryId"))

	_, isExist := r.cateService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	result := r.service.CountByCateID(id)
	response.CreateSuccesResponse(c, &RecipeResponse.RecipeCountByCateResponseDTO{
		Value: result,
	})
}

func (r *recipe) GetByCategoryName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	cateName := converter.MustString(c.Query("name"))

	cates, isExist := r.cateService.FindByName(cateName)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	page := r.service.FindPageByCates(cates, pageable)

	response.CreateSuccesResponse(c, RecipeResponse.CreateRecipeListResponseDTOFromPage(page))
}

func (r *recipe) GetByIngredient(c *gin.Context) {
	pageable := pagination.GetPage(c)
	id := converter.MustUint(c.Param("ingredientId"))

	_, isExist := r.ingreService.FindOneDTO(id)
	if !isExist {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}

	page := r.service.FindPageByIngredientID(id, pageable)

	response.CreateSuccesResponse(c, RecipeResponse.CreateRecipeListResponseDTOFromPage(page))
}

func (r *recipe) GetByIngredientName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	ingredientName := converter.MustString(c.Query("name"))

	ingredientIDs := r.ingreService.FindIDsByName(ingredientName)
	if len(ingredientIDs) == 0 {
		response.CreateErrorResponse(c, "INGREDIENT_NOT_FOUND")
		return
	}

	page := r.service.FindPageByIngredientIDIn(ingredientIDs, pageable)

	response.CreateSuccesResponse(c, RecipeResponse.CreateRecipeListResponseDTOFromPage(page))
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
func (r *recipe) GetByRecipeName(c *gin.Context) {
	pageable := pagination.GetPage(c)
	recipeName := converter.MustString(c.Query("name"))

	page := r.service.FindPageByName(recipeName, pageable)

	response.CreateSuccesResponse(c, RecipeResponse.CreateRecipeListResponseDTOFromPage(page))
}

// Recipe all
// @Summary GetAll
// @Tags PublicRecipe
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeListResponseDTO} "desc"
// @Router /api/public/recipe/getAll [get]
func (r *recipe) GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := r.service.FindPage(pageable)

	response.CreateSuccesResponse(c, RecipeResponse.CreateRecipeListResponseDTOFromPage(page))
}

// Recipe getDetailByID
// @Summary GetDetailByID
// @Tags PublicRecipe
// @Accept json
// @Param id path int true "id"
// @Success 200 {object} response.APIResponseDTO{data=recipe.RecipeDetailResponseDTO} "desc"
// @Router /api/public/recipe/detail/:id [get]
func (r *recipe) GetDetailByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))

	recipeData, isExist := r.service.FindOneWithCate(id)
	if !isExist {
		response.CreateErrorResponse(c, "RECIPE_NOT_FOUND")
		return
	}

	response.CreateSuccesResponse(c, RecipeResponse.CreateRecipeDetailResponseDTO(recipeData))
}
