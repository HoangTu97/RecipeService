package PostResource

import (
	"Food/dto"
	"Food/dto/request"
	PostRequest "Food/dto/request/post"
	"Food/dto/response"
	"Food/helpers/e"
	"Food/helpers/security"
	"Food/service"
	"strings"

	"github.com/gin-gonic/gin"
)

// Post create
// @Summary CreatePost
// @Tags PrivatePost
// @Accept json
// @Security ApiKeyAuth
// @Param body body PostRequest.PostCreateRequestDTO true "body"
// @Success 200 {object} response.APIResponseDTO "desc"
// @Router /api/private/post [post]
func CreatePost(c *gin.Context) {
	userId := security.GetUserID(c)

	var requestDTO PostRequest.PostCreateRequestDTO
	errCode := request.BindAndValid(c, &requestDTO)
	if errCode != e.SUCCESS {
		response.CreateErrorResponse(c, e.GetMsg(errCode))
		return
	}

	userDTO, exists := service.FindOneUserByUserID(userId)
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
	recipeDTO, success = service.SaveRecipe(recipeDTO)
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

	_, success = service.SavePost(postDTO)
	if !success {
		response.CreateErrorResponse(c, "INTERNAL_ERROR")
		return
	}

	response.CreateSuccesResponse(c, nil)
}