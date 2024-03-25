package factory

import "github.com/gin-gonic/gin"

func RouterFactory(root *gin.RouterGroup) {
	group := root.Group("factory")
	{
		group.GET("", GetFactories)
		group.POST("", AddFactory)
		group.PUT("", UpdateFactory)
		group.DELETE(":id", DelFactory)
	}
}
