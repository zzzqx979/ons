package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"ons/controlPanel/http/factory"
	"ons/controlPanel/http/objectModel"
	"ons/controlPanel/http/product"
	"ons/controlPanel/http/tid"
	"ons/dataPanel/http"
	_ "ons/docs"
)

type AIotOption func(*gin.RouterGroup)

var aIotOptions = []AIotOption{
	http.RouterOM,
	tid.RouterTid,
	objectModel.RouterObjectModel,
	product.RouterProduct,
	factory.RouterFactory,
}

// @title			AIot Ons Server
// @version		1.0
// @description	This is AIot Ons server.
func InitAIotHttpRouter() {
	engine := gin.Default()

	rootRouter := engine.Group("ons")
	for _, opt := range aIotOptions {
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
