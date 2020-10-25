package routers

import (
	CategoryResource "Food/controllers/categoryresource"
	CommentResource "Food/controllers/commentresource"
	ImageResource "Food/controllers/imageresource"
	PostResource "Food/controllers/postresource"
	RecipeResource "Food/controllers/reciperesource"
	UserResource "Food/controllers/userresource"
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
			publicCategoryRoutes.GET("/getAll", CategoryResource.GetAll)
		}

		{
			publicRecipeRoutes := publicRoutes.Group("/recipe")
			publicRecipeRoutes.GET("/getAll", RecipeResource.GetAll)
			publicRecipeRoutes.GET("/detail/:id", RecipeResource.GetDetailByID)
			publicRecipeRoutes.GET("/getByCategory/:categoryId", RecipeResource.GetByCategory)
			publicRecipeRoutes.GET("/countByCategory/:categoryId", RecipeResource.GetCountByCategory)
			publicRecipeRoutes.GET("/searchByRecipeName", RecipeResource.GetByRecipeName)
		}

		{
			publicUserRoutes := publicRoutes.Group("/user")
			publicUserRoutes.POST("/register", UserResource.Register)
			publicUserRoutes.POST("/login", UserResource.Login)
		}

		{
			publicImageRoutes := publicRoutes.Group("/image")
			publicImageRoutes.POST("/upload", ImageResource.Upload)
			publicImageRoutes.GET("/:id", ImageResource.FileDisplay)
			publicImageRoutes.GET("/:id/download", ImageResource.Download)
		}

		{
			publicPostRoutes := publicRoutes.Group("/post")
			publicPostRoutes.GET("", PostResource.GetAll)
		}

		{
			publicCommentRoutes := publicRoutes.Group("/comment")
			publicCommentRoutes.GET("/:postId", CommentResource.GetByPostID)
		}
	}

	{
		privateRoutes := apiRoutes.Group("/private")
		privateRoutes.Use(middlewares.Authenticated)
		{
			privatePostRoutes := privateRoutes.Group("/post")
			privatePostRoutes.Use(middlewares.HasAuthority(constants.ROLE_USER))
			privatePostRoutes.POST("", PostResource.CreatePost)
		}
	}
}
