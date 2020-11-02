package controller

import (
	"Food/dto/response"
	CommentResponse "Food/dto/response/comment"
	"Food/helpers/converter"
	"Food/helpers/pagination"
	"Food/service"

	"github.com/gin-gonic/gin"
)

type Comment interface {
	GetByPostID(c *gin.Context)
}

type comment struct {
	service     service.Comment
	postService service.Post
}

func NewComment(service service.Comment, postService service.Post) Comment {
	return &comment{
		service:     service,
		postService: postService,
	}
}

// Comment getByPostID
// @Summary GetByPostID
// @Tags PublicComment
// @Accept json
// @Param postId path uint true "postId"
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {object} response.APIResponseDTO{data=comment.CommentListResponseDTO} "desc"
// @Router /api/public/comment/:postId [get]
func (r *comment) GetByPostID(c *gin.Context) {
	pageable := pagination.GetPage(c)

	id := converter.MustUint(c.Param("postId"))

	_, isExist := r.postService.FindOne(id)
	if !isExist {
		response.CreateErrorResponse(c, "POST_NOT_FOUND")
		return
	}

	page := r.service.FindPageByPostID(id, pageable)

	response.CreateSuccesResponse(c, CommentResponse.CreateCommentListResponseDTOFromPage(page))
}
