package controller

import (
	"Food/dto"
	"Food/dto/request"
	PostRequest "Food/dto/request/post"
	"Food/dto/response"
	PostResponse "Food/dto/response/post"
	"Food/helpers/e"
	"Food/helpers/pagination"
	"Food/helpers/security"
	"Food/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type Post interface {
	GetAll(c *gin.Context)
	CreatePost(c *gin.Context)
}

type post struct {
	service service.Post
	userService service.User
	recipeService service.Recipe
}

func NewPost(service service.Post, userService service.User, recipeService service.Recipe) Post {
	return &post{service: service, userService: userService, recipeService: recipeService}
}

// Post all
// @Summary GetAll
// @Tags PublicPost
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=post.PostListResponseDTO} "desc"
// @Router /api/public/post [get]
func (r *post) GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := r.service.FindPage(pageable)

	response.CreateSuccesResponse(c, PostResponse.CreatePostListResponseDTOFromPage(page))
}

// Post create
// @Summary CreatePost
// @Tags PrivatePost
// @Accept json
// @Security ApiKeyAuth
// @Param body body PostRequest.PostCreateRequestDTO true "body"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/private/post [post]
func (r *post) CreatePost(c *gin.Context) {
	userId := security.GetUserID(c)

	var requestDTO PostRequest.PostCreateRequestDTO
	errCode := request.BindAndValid(c, &requestDTO)
	if errCode != e.SUCCESS {
		response.CreateErrorResponse(c, e.GetMsg(errCode))
		return
	}

	userDTO, exists := r.userService.FindOneByUserID(userId)
	if !exists {
		response.CreateErrorResponse(c, "USER_NOT_FOUND")
		return
	}

	recipeDTO := dto.RecipeDTO{
		Image: requestDTO.RecipeImage,
		Name: requestDTO.RecipeName,
		Duration: requestDTO.RecipeDuration,
		Description: "Nguyên liệu:\n-" + strings.Join(requestDTO.RecipeIngredients, "\n-") + "\nCách làm:\n-" + strings.Join(requestDTO.RecipeSteps, "\n-"),
	}

	var success bool
	recipeDTO, success = r.recipeService.Save(recipeDTO)
	if !success {
		response.CreateErrorResponse(c, "INTERNAL_ERROR")
		return
	}

	postDTO := dto.PostDTO{
		UserID: userDTO.ID,
		Photo: requestDTO.RecipeImage,
		Description: requestDTO.Description,
		HashTags: requestDTO.HashTags,

		RecipeID: recipeDTO.ID,
	}

	_, success = r.service.Save(postDTO)
	if !success {
		response.CreateErrorResponse(c, "INTERNAL_ERROR")
		return
	}

	response.CreateSuccesResponse(c, nil)
}