package upstream

import "github.com/gin-gonic/gin"

func RouterUpstream(root *gin.RouterGroup) {
	upstreamGroup := root.Group("upstream")
	{
		upstreamGroup.GET("", GetUpstream)
		upstreamGroup.POST("", AddUpstream)
		upstreamGroup.PUT("", UpdateUpstream)
		upstreamGroup.DELETE(":id", DelUpstream)
	}
}
