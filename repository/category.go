package repository

import (
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"

	"gorm.io/gorm"
)

type category struct {
	db *gorm.DB
}

type Category interface {
	Save(category models.Category) (models.Category, error)
	FindOne(id uint) (models.Category, error)
	FindOneByName(name string) ([]models.Category, error)
	FindAll() []models.Category
	FindPage(pageable pagination.Pageable) page.Page
	Delete(category models.Category) error
}

func NewCategory(db *gorm.DB) Category {
	return &category{db: db}
}

func (r *category) Save(category models.Category) (models.Category, error) {
	result := r.db.Save(&category)
	if result.Error != nil {
		return category, result.Error
	}
	return category, nil
}

func (r *category) FindOne(id uint) (models.Category, error) {
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

	result := r.db.First(&category, id)
	if result.Error != nil {
		return models.Category{}, result.Error
	}

	// gredis.Set(key, category, 3600)
	return category, nil
}

func (r *category) FindOneByName(name string) ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Where("name LIKE ?", "%" + name + "%").Find(&categories)
	if result.Error != nil {
		return []models.Category{}, result.Error
	}
	return categories, nil
}

func (r *category) FindAll() []models.Category {
	var categories []models.Category
	r.db.Find(&categories)
	return categories
}

func (r *category) FindPage(pageable pagination.Pageable) page.Page {
	var categories []models.Category

	paginator := pagination.Paging(&pagination.Param{
        DB:      r.db.Preload("Recipes"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &categories)

	return page.From(r.toInterfacesFromCategories(categories), paginator.TotalRecord)
}

func (r *category) Delete(category models.Category) error {
	result := r.db.Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *category) toInterfacesFromCategories(categories []models.Category) []interface{} {
	content := make([]interface{}, len(categories))
	for i, v := range categories {
		content[i] = v
	}
	return content
}
