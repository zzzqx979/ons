package tid

import (
	"github.com/gin-gonic/gin"
)

type uriId struct {
	Id int `uri:"id"`
}

func RouterTid(root *gin.RouterGroup) {
	group := root.Group("tid")
	{
		group.GET("", GetTids)
		group.POST("", AddTid)
		group.PUT("", UpdateTid)
		group.DELETE(":id", DelTid)
	}
}
