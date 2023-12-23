package main

import (
	"github.com/gin-gonic/gin"
	// docsのディレクトリを指定
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // ←追記
	"github.com/tatsuya06068/go-gin-redis/docs"
	"github.com/tatsuya06068/go-gin-redis/handlers"
)

// @title gin-swagger todos
// @version 1.0.0
// @description このswaggerはgin-swaggerの見本api
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", handlers.Ping)

	r.Run() // PORT=3000
}
