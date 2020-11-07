package service

import (
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/repository"
	"Food/service/mapper"
)

type Comment interface {
	FindPageByPostID(postID uint, pageable pagination.Pageable) page.Page
}

type comment struct {
	repository repository.Comment
	mapper mapper.Comment
}

func NewComment(repository repository.Comment, mapper mapper.Comment) Comment {
	return &comment{repository: repository, mapper: mapper}
}

func (s *comment) FindPageByPostID(postID uint, pageable pagination.Pageable) page.Page {
	return s.repository.FindPageByPostID(postID, pageable)
}
