package repository

import (
	"Food/config"
	"Food/helpers/page"
	"Food/helpers/pagination"
	"Food/models"
)

func SaveRecipe(recipe models.Recipe) (models.Recipe, error) {
	result := config.GetDB().Save(&recipe)
	if result.Error != nil {
		return recipe, result.Error
	}
	return recipe, nil
}

func FindAllRecipe() []models.Recipe {
	var recipes []models.Recipe
	config.GetDB().Find(&recipes)
	return recipes
}

func FindOneRecipe(id uint) (models.Recipe, error) {
	var recipe models.Recipe
	result := config.GetDB().First(&recipe, id)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil
}

func FindOneRecipePreloadCate(id uint) (models.Recipe, error) {
	var recipe models.Recipe
	result := config.GetDB().Preload("Categories").First(&recipe, id)
	if result.Error != nil {
		return models.Recipe{}, result.Error
	}
	return recipe, nil
}

func FindPageRecipeByCateID(cateID uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Model(&models.Category{ID: cateID}).Joins("Recipes").Joins("Categories").Find(&recipes),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByCates(cates []models.Category, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Model(&cates).Joins("Recipes"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByName(name string, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Where("name LIKE ?", "%" + name + "%").Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByIngredientID(ingredientID uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Where("id IN (?)", config.GetDB().Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id = ?", ingredientID)).Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipeByIngredientIDIn(ingredientIDs []uint, pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Where("id IN (?)", config.GetDB().Table("recipe_ingredients").Select("ingredient_id").Where("recipe_id IN (?)", ingredientIDs)).Preload("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

    return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindPageRecipe(pageable pagination.Pageable) page.Page {
	var recipes []models.Recipe

	paginator := pagination.Paging(&pagination.Param{
        DB:      config.GetDB().Joins("Categories"),
        Page:    pageable.GetPageNumber(),
        Limit:   pageable.GetPageSize(),
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &recipes)

	return page.From(toInterfacesFromRecipes(recipes), paginator.TotalRecord)
}

func FindRecipeByName(name string) []models.Recipe {
	var recipes []models.Recipe
	config.GetDB().Where("name LIKE ?", "%" + name + "%").Find(&recipes)
	return recipes
}

func CountRecipeByCateID(cateID uint) int64 {
	var result int64
	config.GetDB().Table("cate_recipes").Where("category_id = ?", cateID).Count(&result)
	return result
}

func toInterfacesFromRecipes(recipes []models.Recipe) []interface{} {
	content := make([]interface{}, len(recipes))
	for i, v := range recipes {
		content[i] = v
	}
	return content
}
