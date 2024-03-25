package objectModel

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

type DelObjectModelUri struct {
	Id int `uri:"id"`
}

// DeleteObjectModel godoc
// @Summary      删除物模型
// @Tags		 物模型管理(云平台)
// @Accept       json
// @Param        id    path	int  true 	"物模型id"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/object_model/{id} [delete]
func DeleteObjectModel(c *gin.Context) {
	var (
		err error
		uri DelObjectModelUri
	)

	if err = c.BindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if err = db.DelObjectModel(uri.Id); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success"})
}
