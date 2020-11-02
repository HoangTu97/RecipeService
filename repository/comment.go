package repository

import (
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"

	"gorm.io/gorm"
)

type comment struct {
	db *gorm.DB
}

type Comment interface {
	Save(commnent models.Comment) models.Comment
	FindPageByPostID(postID uint, pageable pagination.Pageable) page.Page
}

func NewComment(db *gorm.DB) Comment {
	return &comment{db: db}
}

func (r *comment) Save(commnent models.Comment) models.Comment {
	r.db.Create(&commnent)
	return commnent
}

func (r *comment) FindPageByPostID(postID uint, pageable pagination.Pageable) page.Page {
	var comments []models.Comment

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.db.Where("post_id = ?", postID).Find(&comments).Preload("User"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &comments)

    return page.From(r.toInterfacesFromComments(comments), paginator.TotalRecord)
}

func (r *comment) toInterfacesFromComments(comments []models.Comment) []interface{} {
	content := make([]interface{}, len(comments))
	for i, v := range comments {
		content[i] = v
	}
	return content
}
