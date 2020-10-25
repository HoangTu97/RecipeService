package commentresource

import (
	"Food/dto/response"
	"Food/dto/response/comment"
	"Food/helpers/converter"
	"Food/helpers/pagination"
	"Food/service"

	"github.com/gin-gonic/gin"
)

// Comment getByPostID
// @Summary GetByPostID
// @Tags PublicComment
// @Accept json
// @Param postId path uint true "postId"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=comment.CommentListResponseDTO} "desc"
// @Router /api/public/comment/:postId [get]
func GetByPostID(c *gin.Context) {
	pageable := pagination.GetPage(c)

	id := converter.MustUint(c.Param("postId"))

	_, isExist := service.FindOnePost(id)
	if !isExist {
		response.CreateErrorResponse(c, "POST_NOT_FOUND")
		return
	}

	page := service.FindPageCommentByPostID(id, pageable)

	response.CreateSuccesResponse(c, comment.CreateCommentListResponseDTOFromPage(page))
}