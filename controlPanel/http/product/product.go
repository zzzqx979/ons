package product

import "github.com/gin-gonic/gin"

func RouterProduct(root *gin.RouterGroup) {
	group := root.Group("product")
	{
		group.GET("", GetProducts)
		group.POST("", AddProduct)
		group.PUT("", UpdateProduct)
		group.DELETE(":id", DelProduct)
	}
}
