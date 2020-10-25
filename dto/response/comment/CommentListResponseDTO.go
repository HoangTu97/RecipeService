package comment

import (
	"Food/helpers/page"
	"Food/models"
	"time"
)

// CommentListResponseDTO godoc
type CommentListResponseDTO struct {
	Items      []CommentListItemResponseDTO `json:"items"`
	TotalItems int64                        `json:"totalItems"`
}

// CommentListItemResponseDTO godoc
type CommentListItemResponseDTO struct {
	ID uint `json:"id"`

	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`

	Text      string    `json:"text"`
	CountLike uint      `json:"count_like"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateCommentListResponseDTOFromPage create page from page models.Comment
func CreateCommentListResponseDTOFromPage(page page.Page) *CommentListResponseDTO {
	result := make([]CommentListItemResponseDTO, page.GetTotalElements())
	for i, v := range page.GetContent() {
		entity := v.(models.Comment)
		user := entity.User

		result[i] = CommentListItemResponseDTO{
			ID: entity.ID,

			UserID:   user.ID,
			UserName: user.Name,

			Text:      entity.Description,
			CreatedAt: entity.CreatedAt,
			CountLike: 0,
		}
	}
	return &CommentListResponseDTO{
		Items:      result,
		TotalItems: page.GetTotalItems(),
	}
}
