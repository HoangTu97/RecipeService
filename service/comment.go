package service

import (
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/repository"
)

func FindPageCommentByPostID(postID uint, pageable pagination.Pageable) page.Page {
	return repository.FindPageCommentByPostID(postID, pageable)
}
