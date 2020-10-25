package repository

import (
	"Food/config"
	"Food/helpers/converter"
	"Food/helpers/gredis"
	"Food/helpers/logging"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
	"encoding/json"
)

func SavePost(post models.Post) (models.Post, error) {
	result := config.GetDB().Save(&post)
	if result.Error != nil {
		return post, result.Error
	}
	return post, nil
}

func FindOnePost(id uint) (models.Post, error) {
	var post models.Post

	key := "POST_" + converter.ToStr(id)
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err == nil {
			logging.Info("FindOneCate", err)
			return models.Post{}, err
		}
		_ = json.Unmarshal(data, &post)
		return post, nil
	}

	result := config.GetDB().First(&post, id)
	if result.Error != nil {
		return models.Post{}, result.Error
	}

	_ = gredis.Set(key, post, 3600)
	return post, nil
}

func FindPagePost(pageable pagination.Pageable) page.Page {
	var posts []models.Post

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Joins("User").Joins("Recipe"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        ShowSQL: true,
	}, &posts)

	return page.From(toInterfacesFromPost(posts), paginator.TotalRecord)
}

func toInterfacesFromPost(posts []models.Post) []interface{} {
	content := make([]interface{}, len(posts))
	for i, v := range posts {
		content[i] = v
	}
	return content
}