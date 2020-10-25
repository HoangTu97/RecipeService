package post

import (
	"Food/helpers/page"
	"Food/models"
	"time"
)

// PostListResponseDTO godoc
type PostListResponseDTO struct {
	Items      []PostListItemResponseDTO `json:"items"`
	TotalItems int64                     `json:"totalItems"`
}

// PostListItemResponseDTO godoc
type PostListItemResponseDTO struct {
	ID uint `json:"id"`

	UserID    uint      `json:"user_id"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`

	RecipeID    uint     `json:"recipe_id"`
	RecipePhoto string   `json:"recipe_photo"`
	RecipeTitle string   `json:"recipe_title"`
	RecipeDesc  string   `json:"recipe_desc"`
	HashTags    []string `json:"tags"`

	CountLike uint `json:"count_like"`
}

// CreatePostListResponseDTOFromPage create page from page models.Post
func CreatePostListResponseDTOFromPage(page page.Page) *PostListResponseDTO {
	result := make([]PostListItemResponseDTO, page.GetTotalElements())
	for i, v := range page.GetContent() {
		entity := v.(models.Post)
		user := entity.User
		recipe := entity.Recipe

		result[i] = PostListItemResponseDTO{
			ID: entity.ID,

			UserID:    user.ID,
			UserName:  user.Name,
			CreatedAt: entity.CreatedAt,

			RecipeID:    recipe.ID,
			RecipePhoto: recipe.Image,
			RecipeTitle: recipe.Name,
			RecipeDesc:  recipe.Description,
			// HashTags: nil,

			CountLike: 0,
		}
	}
	return &PostListResponseDTO{
		Items:      result,
		TotalItems: page.GetTotalItems(),
	}
}
