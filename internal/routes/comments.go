package routes

import (
	"github.com/gin-gonic/gin"
	"gorum/internal/auth"
	"gorum/internal/controllers"
)

func CommentRoutes(r *gin.Engine, c *controllers.MainController) {
	commentRoutes := r.Group("/comments")

	commentRoutes.PUT("/:comment_id", auth.Handler, c.UpdateComment)
	commentRoutes.DELETE("/:comment_id", auth.Handler, c.DeleteComment)
}
