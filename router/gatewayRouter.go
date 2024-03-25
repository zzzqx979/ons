package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"ons/controlPanel/http/upstream"
	"ons/dataPanel/http"
)

type GatewayOptions func(group *gin.RouterGroup)

var gatewayOptions = []GatewayOptions{
	http.RouterOM,
	upstream.RouterUpstream,
}

// @title			Gateway Ons Server
// @version		1.0
// @description	This is Gateway Ons server.
func InitGatewayHttpServer() {
	engine := gin.Default()

	rootRouter := engine.Group("/ons")
	for _, opt := range gatewayOptions {
		opt(rootRouter)
	}

	// use ginSwagger middleware to serve the API docs
	rootRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// health check
	rootRouter.GET("/health", func(context *gin.Context) {
		return
	})
	err := engine.Run(":" + viper.GetString("server.port"))
	if err != nil {
		panic(fmt.Errorf("fatal init http router errors:%w", err))
	}
}
