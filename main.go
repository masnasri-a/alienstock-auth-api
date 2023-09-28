package main

import (
	"alienstock-auth-api/config"
	"alienstock-auth-api/controller"
	"alienstock-auth-api/docs"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	fmt.Println("Load a env file")
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}
	config.Ping()
}

func main() {
	route := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	basePath := "/api/v1"
	apiV1 := route.Group(basePath)
	controller.AuthRouter(apiV1)
	route.Run()
}
