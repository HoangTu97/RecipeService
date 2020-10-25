package service

import (
	"Food/dto"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/repository"
	"Food/service/mapper"
)

func SavePost(postDTO dto.PostDTO) (dto.PostDTO, bool) {
	post := mapper.ToPost(postDTO)
	var err error
	post, err = repository.SavePost(post)
	if err != nil {
		return postDTO, false
	}
	return mapper.ToPostDTO(post), true
}

func FindOnePost(id uint) (dto.PostDTO, bool) {
	post, err := repository.FindOnePost(id)
	if err != nil {
		return dto.PostDTO{}, false
	}
	return mapper.ToPostDTO(post), true
}

func FindPagePost(pageable pagination.Pageable) page.Page {
	return repository.FindPagePost(pageable)
}