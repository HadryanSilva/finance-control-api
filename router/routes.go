package router

import (
	"github.com/HadryanSilva/finance-control-api/docs"
	"github.com/HadryanSilva/finance-control-api/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	accountRoutes := router.Group(basePath)
	{
		accountRoutes.GET("/account", handler.GetAccountHandler)
		accountRoutes.POST("/account", handler.SaveAccountHandler)
		accountRoutes.DELETE("/account")
		accountRoutes.PUT("/account", handler.UpdateAccountHandler)
		accountRoutes.GET("/account/list")
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
