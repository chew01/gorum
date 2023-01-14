package main

import (
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorum/docs"
	"gorum/internal/router"
)

func main() {
	r := router.Setup()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
