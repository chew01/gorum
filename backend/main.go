package main

import (
	"fmt"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorum/docs"
	"gorum/internal/controllers"
	"gorum/internal/database"
	"gorum/internal/router"
)

func main() {
	db, err := database.New("./root.db")

	if err != nil {
		fmt.Println("Error opening database connection:", err)
		return
	}

	if err := db.Init(); err != nil {
		fmt.Println("Error initializing database: ", err)
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			fmt.Println("Error closing database connection:", err)
			return
		}
	}()

	c := controllers.NewMainController(db)
	r := router.Setup(c)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	_ = r.Run(":8080")
}
