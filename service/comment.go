package service

import (
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/repository"
)

type Comment interface {
	FindPageByPostID(postID uint, pageable pagination.Pageable) page.Page
}

type comment struct {
	repository repository.Comment
}

func NewComment(repository repository.Comment) Comment {
	return &comment{repository: repository}
}

func (s *comment) FindPageByPostID(postID uint, pageable pagination.Pageable) page.Page {
	return s.repository.FindPageByPostID(postID, pageable)
}
