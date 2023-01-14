package routes

import (
	"github.com/gin-gonic/gin"
	"gorum/internal/controllers"
)

func CommentRoutes(r *gin.Engine) {
	commentRoutes := r.Group("/comments")

	commentRoutes.PUT("/:comment_id", controllers.UpdateComment)
	commentRoutes.DELETE("/:comment_id", controllers.DeleteComment)
}
