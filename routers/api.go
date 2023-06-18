package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	//"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)

	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/create", controllers.CreatePhoto)
		photoRouter.GET("/photo", controllers.GetAllPhoto)
		photoRouter.GET("/:photoId", controllers.GetOnePhoto)
		photoRouter.GET("/photoUser", controllers.GetPhotos)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)

	}

	socialmediaRouter := router.Group("/socmeds")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.POST("/create", controllers.CreateSocialmedia)
		socialmediaRouter.GET("/social", controllers.GetAllSocialMedia)
		socialmediaRouter.GET("/:socialMediaId", controllers.GetOneSocialMedia)
		socialmediaRouter.GET("/socialMediaUser", controllers.GetSocialMedias)
		socialmediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialmediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)

	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/create", controllers.CreateComment)
		commentRouter.GET("/get", controllers.GetAllComments)
		commentRouter.GET("/:commentId", controllers.GetOneComment)
		commentRouter.GET("/getCommentUser", controllers.GetComments)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)

	}

	return router
}
