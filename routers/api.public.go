package routers

import (
	"Food/config"

	"github.com/gin-gonic/gin"
)

func registerPublicApi(apiRoutes *gin.RouterGroup) {
	publicRoutes := apiRoutes.Group("/public")
	// Api declare
	{
		publicCategoryRoutes := publicRoutes.Group("/category")
		publicCategoryRoutes.GET("/getAll", config.CateController.GetAll)
	}

	{
		publicRecipeRoutes := publicRoutes.Group("/recipe")
		publicRecipeRoutes.GET("/getAll", config.RecipeController.GetAll)
		publicRecipeRoutes.GET("/detail/:id", config.RecipeController.GetDetailByID)
		publicRecipeRoutes.GET("/getByCategory/:categoryId", config.RecipeController.GetByCategory)
		publicRecipeRoutes.GET("/countByCategory/:categoryId", config.RecipeController.GetCountByCategory)
		publicRecipeRoutes.GET("/searchByRecipeName", config.RecipeController.GetByRecipeName)
	}

	{
		publicUserRoutes := publicRoutes.Group("/user")
		publicUserRoutes.POST("/register", config.UserController.Register)
		publicUserRoutes.POST("/login", config.UserController.Login)
	}

	{
		publicImageRoutes := publicRoutes.Group("/image")
		publicImageRoutes.POST("/upload", config.ImageController.Upload)
		publicImageRoutes.GET("/:id", config.ImageController.FileDisplay)
		publicImageRoutes.GET("/:id/download", config.ImageController.Download)
	}

	{
		publicPostRoutes := publicRoutes.Group("/post")
		publicPostRoutes.GET("", config.PostController.GetAll)
	}

	{
		publicCommentRoutes := publicRoutes.Group("/comment")
		publicCommentRoutes.GET("/:postId", config.CommentController.GetByPostID)
	}
	// Api declare end : dont remove
}
