package repository

import (
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"

	"gorm.io/gorm"
)

type recipe struct {
	db *gorm.DB
}

type Recipe interface {
	Save(recipe models.Recipe) (models.Recipe, error)
	FindAll() []models.Recipe
	FindOne(id uint) (models.Recipe, error)
	FindOnePreloadCate(id uint) (models.Recipe, error)
	FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page
	FindPageByCates(cates []models.Category, pageable pagination.Pageable) page.Page
	FindPageByName(name string, pageable pagination.Pageable) page.Page
	FindPageByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page
	FindPageByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page
	FindPage(pageable pagination.Pageable) page.Page
	FindByName(name string) []models.Recipe
	CountByCateID(cateID uint) int64
}

func NewRecipe(db *gorm.DB) Recipe {
	return &recipe{db: db}
}

func (r *recipe) Save(recipe models.Recipe) (models.Recipe, error) {
	result := r.db.Save(&recipe)
	if result.Error != nil {
		return recipe, result.Error
	}
	return recipe, nil
}

func (r *recipe) FindAll() []models.Recipe {
	var recipes []models.Recipe
	r.db.Find(&recipes)
	return recipes
}

func (r *recipe) FindOne(id uint) (models.Recipe, error) {
	var recipe models.Recipe
	result := r.db.First(&recipe, id)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil
}

func (r *recipe) FindOnePreloadCate(id uint) (models.Recipe, error) {
	var recipe models.Recipe
	result := r.db.Preload("Categories").First(&recipe, id)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil
}

func (r *recipe) FindPageByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
		DB:      r.db.Model(&models.Category{ID: cateID}).Joins("Recipes").Joins("Categories").Find(&recipes),
		Page:    pageable.GetPageNumber(),
		Limit:   pageable.GetPageSize(),
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &recipes)

	return page.From(r.toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipe) FindPageByCates(cates []models.Category, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
		DB:      r.db.Model(&cates).Joins("Recipes"),
		Page:    pageable.GetPageNumber(),
		Limit:   pageable.GetPageSize(),
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &recipes)

	return page.From(r.toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipe) FindPageByName(name string, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
		DB:      r.db.Where("name LIKE ?", "%"+name+"%").Preload("Categories"),
		Page:    pageable.GetPageNumber(),
		Limit:   pageable.GetPageSize(),
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &recipes)

	return page.From(r.toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipe) FindPageByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
		DB:      r.db.Where("id IN (?)", r.db.Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id = ?", ingredientID)).Preload("Categories"),
		Page:    pageable.GetPageNumber(),
		Limit:   pageable.GetPageSize(),
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &recipes)

	return page.From(r.toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipe) FindPageByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
		DB:      r.db.Where("id IN (?)", r.db.Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id IN (?)", ingredientIDs)).Preload("Categories"),
		Page:    pageable.GetPageNumber(),
		Limit:   pageable.GetPageSize(),
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &recipes)

	return page.From(r.toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipe) FindPage(pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
		DB:      r.db.Joins("Categories"),
		Page:    pageable.GetPageNumber(),
		Limit:   pageable.GetPageSize(),
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &recipes)

	return page.From(r.toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func (r *recipe) FindByName(name string) []models.Recipe {
	var recipes []models.Recipe
	r.db.Where("name LIKE ?", "%"+name+"%").Find(&recipes)
	return recipes
}

func (r *recipe) CountByCateID(cateID uint) int64 {
	var result int64
	r.db.Table("cate_recipes").Where("category_id = ?", cateID).Count(&result)
	return result
}

func (r *recipe) toInterfacesFromRecipes(recipes []models.Recipe) []interface{} {
	content := make([]interface{}, len(recipes))
	for i, v := range recipes {
		content[i] = v
	}
	return content
}
