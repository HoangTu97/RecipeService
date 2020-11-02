package controller

import (
	"Food/dto/response"
	CateResponse "Food/dto/response/category"
	"Food/helpers/converter"
	"Food/helpers/pagination"
	"Food/service"

	"github.com/gin-gonic/gin"
)

type Category interface {
	GetAll(c *gin.Context)
	GetAllMini(c *gin.Context)
	GetByID(c *gin.Context)
	GetNameByID(c *gin.Context)
}

type category struct {
	service service.Category
}

func NewCategory(service service.Category) Category {
	return &category{service: service}
}

// Category all
// @Summary GetAll
// @Tags PublicCategory
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=category.CategoryListResponseDTO} "desc"
// @Router /api/public/category/getAll [get]
func (r *category) GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := r.service.FindPage(pageable)

	response.CreateSuccesResponse(c, CateResponse.CreateCategoryListResponseDTOFromPage(page))
}

// GetAllMini all mini
// @Summary GetAllMini
// @Tags PublicCategory
// @Accept json
// @Success 200 {object} response.APIResponseDTO{data=category.CategoryMiniListResponseDTO} "desc"
// @Router /api/public/category/getAll [get]
func (r *category) GetAllMini(c *gin.Context) {
	categoryDTOS := r.service.FindAll()

	response.CreateSuccesResponse(c, categoryDTOS)
}

func (r *category) GetByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))

	categoryDTO, isExist := r.service.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	response.CreateSuccesResponse(c, &CateResponse.CategoryDetailResponseDTO{
		ID:    categoryDTO.ID,
		Name:  categoryDTO.Name,
		Image: categoryDTO.Image,
	})
}

func (r *category) GetNameByID(c *gin.Context) {
	id := converter.MustUint(c.Param("id"))

	categoryDTO, isExist := r.service.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "CATEGORY_NOT_FOUND")
		return
	}

	response.CreateSuccesResponse(c, &CateResponse.CategoryNameResponseDTO{
		Name: categoryDTO.Name,
	})
}
