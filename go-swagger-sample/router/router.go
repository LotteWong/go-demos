package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-swagger-sample/controller"
	"go-swagger-sample/docs"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userGroup := router.Group("/api/v1/users")
	{
		userGroup.GET("", controller.ListUsers)
		userGroup.POST("", controller.CreateUser)
		userGroup.GET("/:id", controller.GetUser)
	}

	return router
}
