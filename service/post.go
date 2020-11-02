package service

import (
	"Food/dto"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/repository"
	"Food/service/mapper"
)

type Post interface {
	Save(postDTO dto.PostDTO) (dto.PostDTO, bool)
	FindOne(id uint) (dto.PostDTO, bool)
	FindPage(pageable pagination.Pageable) page.Page
}

type post struct {
	repository repository.Post
}

func NewPost(repository repository.Post) Post {
	return &post{repository: repository}
}

func (s *post) Save(postDTO dto.PostDTO) (dto.PostDTO, bool) {
	post := mapper.ToPost(postDTO)
	var err error
	post, err = s.repository.Save(post)
	if err != nil {
		return postDTO, false
	}
	return mapper.ToPostDTO(post), true
}

func (s *post) FindOne(id uint) (dto.PostDTO, bool) {
	post, err := s.repository.FindOne(id)
	if err != nil {
		return dto.PostDTO{}, false
	}
	return mapper.ToPostDTO(post), true
}

func (s *post) FindPage(pageable pagination.Pageable) page.Page {
	return s.repository.FindPage(pageable)
}