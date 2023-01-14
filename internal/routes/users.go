package routes

import "github.com/gin-gonic/gin"
import "gorum/internal/controllers"

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")

	userRoutes.GET("/", controllers.GetUser)
	userRoutes.POST("/", controllers.LoginUser)
	userRoutes.POST("/create", controllers.CreateUser)
}
