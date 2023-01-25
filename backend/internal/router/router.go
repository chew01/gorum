package router

import (
	"github.com/gin-gonic/gin"
	"gorum/internal/controllers"
)
import "gorum/internal/routes"

func Setup(c *controllers.MainController) *gin.Engine {
	r := gin.Default()
	setupRoutes(r, c)
	return r
}

func setupRoutes(r *gin.Engine, c *controllers.MainController) {
	routes.UserRoutes(r, c)
	routes.PostRoutes(r, c)
	routes.CommentRoutes(r, c)
}
