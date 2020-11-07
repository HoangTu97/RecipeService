package config

import (
	"Food/controller"
	"Food/helpers/cache"
	"Food/repository"
	"Food/service"
	"Food/service/mapper"

	"gorm.io/gorm"
)

var (
	CateController controller.Category
	CommentController controller.Comment
	ImageController controller.Image
	IngredientController controller.Ingredient
	PostController controller.Post
	RecipeController controller.Recipe
	UserController controller.User
)

func SetupController(db *gorm.DB) {
	cache := cache.NewRedis(*RedisSetting)

	cateMapper := mapper.NewCategory()
	commentMapper := mapper.NewComment()
	ingredientMapper := mapper.NewIngredient()
	postMapper := mapper.NewPost()
	recipeIngreMapper := mapper.NewRecipeIngredient()
	recipeMapper := mapper.NewRecipe()
	userMapper := mapper.NewUser()

	cateRepo := repository.NewCategory(db)
	commentRepo := repository.NewComment(db)
	ingreRepo := repository.NewIngredient(db)
	postRepo := repository.NewPost(db)
	recipeIngreRepo := repository.NewRecipeIngredients(db)
	recipeRepo := repository.NewRecipe(db)
	// userRecipeInteractRepo := repository.NewUserRecipeInteraction(db)
	userRepo := repository.NewUser(db)

	cateService := service.NewCategory(cateRepo, cateMapper)
	commentService := service.NewComment(commentRepo, commentMapper)
	imageService := service.NewImage()
	ingreService := service.NewIngredient(ingreRepo, ingredientMapper)
	postService := service.NewPost(postRepo, postMapper)
	recipeIngreService := service.NewRecipeIngredients(recipeIngreRepo, recipeIngreMapper)
	recipeService := service.NewRecipe(recipeRepo, recipeMapper)
	// userRecipeInteractService := service.NewUser(userRecipeInteractRepo)
	userService := service.NewUser(userRepo, userMapper)

	userServiceProxy := service.NewUserProxy(userService, cache)

	CateController = controller.NewCategory(cateService)
	CommentController = controller.NewComment(commentService, postService)
	ImageController = controller.NewImage(imageService)
	IngredientController = controller.NewIngredient(ingreService, recipeService, recipeIngreService)
	PostController = controller.NewPost(postService, userService, recipeService)
	RecipeController = controller.NewRecipe(recipeService, cateService, ingreService)
	UserController = controller.NewUser(userServiceProxy)
}
