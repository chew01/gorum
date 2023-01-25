package routes

import (
	"github.com/gin-gonic/gin"
	"gorum/internal/auth"
)
import "gorum/internal/controllers"

func UserRoutes(r *gin.Engine, c *controllers.MainController) {
	userRoutes := r.Group("/auth")

	userRoutes.GET("/", auth.Handler, c.GetUser)
	userRoutes.POST("/", c.LoginUser)
	userRoutes.POST("/create", c.CreateUser)
}
