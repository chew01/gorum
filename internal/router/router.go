package router

import "github.com/gin-gonic/gin"
import "gorum/internal/routes"

func Setup() *gin.Engine {
	r := gin.Default()
	setupRoutes(r)
	return r
}

func setupRoutes(r *gin.Engine) {
	routes.UserRoutes(r)
	routes.PostRoutes(r)
	routes.CommentRoutes(r)
}
