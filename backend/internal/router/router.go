package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorum/internal/controllers"
)
import "gorum/internal/routes"

func Setup(c *controllers.MainController) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers", "Authorization"},
	}))
	setupRoutes(r, c)
	return r
}

func setupRoutes(r *gin.Engine, c *controllers.MainController) {
	routes.UserRoutes(r, c)
	routes.PostRoutes(r, c)
	routes.CommentRoutes(r, c)
}
