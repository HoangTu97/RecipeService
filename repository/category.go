package repository

import (
	"Food/config"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
)

func SaveCate(category models.Category) (models.Category, error) {
	result := config.GetDB().Save(&category)
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

func FindOneCate(id uint) (models.Category, error) {
	var category models.Category

	// key := "CATE_" + converter.ToStr(id)
	// if gredis.Exists(key) {
	// 	data, err := gredis.Get(key)
	// 	if err == nil {
	// 		json.Unmarshal(data, &category)
	// 		return category, nil
	// 	}
	// 	logging.Info("FindOneCate", err)
	// }

	result := config.GetDB().First(&category, id)
	if result.Error != nil {
		return models.Category{}, result.Error
	}

	// gredis.Set(key, category, 3600)
	return category, nil
}

func FindOneCateByName(name string) ([]models.Category, error) {
	var categories []models.Category
	result := config.GetDB().Where("name LIKE ?", "%" + name + "%").Find(&categories)
	if result.Error != nil {
		return []models.Category{}, result.Error
	}
	return categories, nil
}

func FindAllCate() []models.Category {
	var categories []models.Category
	config.GetDB().Find(&categories)
	return categories
}

func FindPageCate(pageable pagination.Pageable) page.Page {
	var categories []models.Category

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Preload("Recipes"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &categories)

	return page.From(toInterfacesFromCategories(categories), paginator.TotalRecord)
}

func DeleteCate(category models.Category) error {
	result := config.GetDB().Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func toInterfacesFromCategories(categories []models.Category) []interface{} {
	content := make([]interface{}, len(categories))
	for i, v := range categories {
		content[i] = v
	}
	return content
}
