package config

import (
  "Food/controller"
  "Food/helpers/jwt"
  "Food/pkg/cache"
  "Food/repository"
  "Food/repository/impl"
  "Food/repository/proxy"
  "Food/service"
  "Food/service/mapper"

  "gorm.io/gorm"
)

var (
  // Controllers globale declare
  CateController controller.Category
  CommentController controller.Comment
  ImageController controller.Image
  IngredientController controller.Ingredient
  PostController controller.Post
  RecipeController controller.Recipe
  UserController controller.User
  // Controllers globale declare end : dont remove
)

func SetupController(db *gorm.DB, jwtManager jwt.JwtManager, cache cache.Cache) {
  // Mappers declare
  cateMapper := mapper.NewCategory()
  commentMapper := mapper.NewComment()
  ingredientMapper := mapper.NewIngredient()
  postMapper := mapper.NewPost()
  recipeIngreMapper := mapper.NewRecipeIngredient()
  recipeMapper := mapper.NewRecipe()
  userMapper := mapper.NewUser()
  // Mappers declare end : dont remove

  // Repositories declare
  cateRepo := repository.NewCategory(db)
  commentRepo := repository.NewComment(db)
  ingreRepo := repository.NewIngredient(db)
  postRepo := repository.NewPost(db)
  recipeIngreRepo := repository.NewRecipeIngredients(db)
  recipeRepo := repository.NewRecipe(db)
  // userRecipeInteractRepo := repository.NewUserRecipeInteraction(db)
  userRepo := repository_impl.NewUser(db)
  // Repositories declare end : dont remove

  // Proxy Repositories declare
  userRepoProxy := repository_proxy.NewUser(userRepo)
  // Proxy Repositories declare end : dont remove

  // Services declare
  cateService := service.NewCategory(cateRepo, cateMapper)
  commentService := service.NewComment(commentRepo, commentMapper)
  imageService := service.NewImage()
  ingreService := service.NewIngredient(ingreRepo, ingredientMapper)
  postService := service.NewPost(postRepo, postMapper)
  recipeIngreService := service.NewRecipeIngredients(recipeIngreRepo, recipeIngreMapper)
  recipeService := service.NewRecipe(recipeRepo, recipeMapper)
  // userRecipeInteractService := service.NewUser(userRecipeInteractRepo)
  userService := service.NewUser(userRepoProxy, userMapper, jwtManager)
  // Services declare end : dont remove

  // Proxy Services declare
  userServiceProxy := service.NewUserProxy(userService, cache)
  // Proxy Services declare end : dont remove

  // Controllers declare
  CateController = controller.NewCategory(cateService)
  CommentController = controller.NewComment(commentService, postService)
  ImageController = controller.NewImage(imageService)
  IngredientController = controller.NewIngredient(ingreService, recipeService, recipeIngreService)
  PostController = controller.NewPost(postService, userServiceProxy, recipeService)
  RecipeController = controller.NewRecipe(recipeService, cateService, ingreService)
  UserController = controller.NewUser(userServiceProxy)
  // Controllers declare end : dont remove
}
