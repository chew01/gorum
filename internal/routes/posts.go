package routes

import (
	"github.com/gin-gonic/gin"
	"gorum/internal/auth"
	"gorum/internal/controllers"
)

func PostRoutes(r *gin.Engine, c *controllers.MainController) {
	postRoutes := r.Group("/posts")

	postRoutes.GET("/", c.GetPosts)
	postRoutes.POST("/", auth.Handler, c.CreatePost)

	postRoutes.GET("/:post_id", c.GetPost)
	postRoutes.PUT("/:post_id", auth.Handler, c.UpdatePost)
	postRoutes.DELETE("/:post_id", auth.Handler, c.DeletePost)

	postRoutes.POST("/:post_id/comments", auth.Handler, c.CreateComment)
}
