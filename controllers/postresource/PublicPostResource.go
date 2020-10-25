package PostResource

import (
	"Food/dto/response"
	"Food/dto/response/post"
	"Food/helpers/pagination"
	"Food/service"

	"github.com/gin-gonic/gin"
)

// Post all
// @Summary GetAll
// @Tags PublicPost
// @Accept json
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=post.PostListResponseDTO} "desc"
// @Router /api/public/post [get]
func GetAll(c *gin.Context) {
	pageable := pagination.GetPage(c)

	page := service.FindPagePost(pageable)

	response.CreateSuccesResponse(c, post.CreatePostListResponseDTOFromPage(page))
}