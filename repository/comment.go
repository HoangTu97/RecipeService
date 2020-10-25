package repository

import (
	"Food/config"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
)

func SaveComment(commnent models.Comment) models.Comment {
	config.GetDB().Create(&commnent)
	return commnent
}

func FindPageCommentByPostID(postID uint, pageable pagination.Pageable) page.Page {
	var comments []models.Comment

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Where("post_id = ?", postID).Find(&comments).Preload("User"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &comments)

    return page.From(toInterfacesFromComments(comments), paginator.TotalRecord)
}

func toInterfacesFromComments(comments []models.Comment) []interface{} {
	content := make([]interface{}, len(comments))
	for i, v := range comments {
		content[i] = v
	}
	return content
}
