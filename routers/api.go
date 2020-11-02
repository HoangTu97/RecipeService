package routers

import (
	"Food/config"
	"Food/helpers/constants"
	"Food/middlewares"

	"github.com/gin-gonic/gin"
)

// InitRouterApi InitRouterApi
func InitRouterApi(r *gin.Engine) {
	apiRoutes := r.Group("/api")
	{
		publicRoutes := apiRoutes.Group("/public")
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
	}

	{
		privateRoutes := apiRoutes.Group("/private")
		privateRoutes.Use(middlewares.Authenticated)
		{
			privatePostRoutes := privateRoutes.Group("/post")
			privatePostRoutes.Use(middlewares.HasAuthority(constants.ROLE_USER))
			privatePostRoutes.POST("", config.PostController.CreatePost)
		}
	}
}
