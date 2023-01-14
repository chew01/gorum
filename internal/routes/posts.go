package routes

import (
	"github.com/gin-gonic/gin"
	"gorum/internal/controllers"
)

func PostRoutes(r *gin.Engine) {
	postRoutes := r.Group("/posts")

	postRoutes.GET("/", controllers.GetPosts)
	postRoutes.POST("/", controllers.CreatePost)

	// postRoutes.GET("/:post_id", controllers.GetPostContent)
	postRoutes.PUT("/:post_id", controllers.UpdatePost)
	postRoutes.DELETE("/:post_id", controllers.DeletePost)

	// postRoutes.GET("/:post_id/comments", controllers.GetComments)
	postRoutes.POST("/:post_id/comments", controllers.CreateComment)
}
