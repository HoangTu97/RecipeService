package config

import (
	"Food/controller"
	"Food/repository"
	"Food/service"
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

func SetupDI() {
	db := GetDB()

	cateRepo := repository.NewCategory(db)
	commentRepo := repository.NewComment(db)
	ingreRepo := repository.NewIngredient(db)
	postRepo := repository.NewPost(db)
	recipeIngreRepo := repository.NewRecipeIngredients(db)
	recipeRepo := repository.NewRecipe(db)
	// userRecipeInteractRepo := repository.NewUserRecipeInteraction(db)
	userRepo := repository.NewUser(db)

	cateService := service.NewCategory(cateRepo)
	commentService := service.NewComment(commentRepo)
	ingreService := service.NewIngredient(ingreRepo)
	postService := service.NewPost(postRepo)
	recipeIngreService := service.NewRecipeIngredients(recipeIngreRepo)
	recipeService := service.NewRecipe(recipeRepo)
	// userRecipeInteractService := service.NewUser(userRecipeInteractRepo)
	userService := service.NewUser(userRepo)

	CateController = controller.NewCategory(cateService)
	CommentController = controller.NewComment(commentService, postService)
	ImageController = controller.NewImage()
	IngredientController = controller.NewIngredient(ingreService, recipeService, recipeIngreService)
	PostController = controller.NewPost(postService, userService, recipeService)
	RecipeController = controller.NewRecipe(recipeService, cateService, ingreService)
	UserController = controller.NewUser(userService)
}
