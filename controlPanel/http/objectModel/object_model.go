package objectModel

import "github.com/gin-gonic/gin"

func RouterObjectModel(root *gin.RouterGroup) {
	{
		groupObjectModel := root.Group("object_model")
		{
			groupObjectModel.GET("", GetObjectModels)
			groupObjectModel.POST("", AddObjectModel)
			groupObjectModel.PUT("", UpdateObjectModel)
			groupObjectModel.DELETE(":id", DeleteObjectModel)
		}
	}
}
