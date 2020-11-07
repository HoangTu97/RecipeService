package repository

import (
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"

	"gorm.io/gorm"
)

type post struct {
	db *gorm.DB
}

type Post interface {
	Save(post models.Post) (models.Post, error)
	FindOne(id uint) (models.Post, error)
	FindPage(pageable pagination.Pageable) page.Page
}

func NewPost(db *gorm.DB) Post {
	return &post{db: db}
}

func (r *post) Save(post models.Post) (models.Post, error) {
	result := r.db.Save(&post)
	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

func (r *post) FindOne(id uint) (models.Post, error) {
	var post models.Post

	result := r.db.First(&post, id)
	if result.Error != nil {
		return models.Post{}, result.Error
	}

	return post, nil
}

func (r *post) FindPage(pageable pagination.Pageable) page.Page {
	var posts []models.Post

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.db.Joins("User").Joins("Recipe"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        ShowSQL: true,
	}, &posts)

	return page.From(r.toInterfacesFromPost(posts), paginator.TotalRecord)
}

func (r *post) toInterfacesFromPost(posts []models.Post) []interface{} {
	content := make([]interface{}, len(posts))
	for i, v := range posts {
		content[i] = v
	}
	return content
}