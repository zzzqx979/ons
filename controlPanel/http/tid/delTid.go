package tid

import (
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"

	"github.com/gin-gonic/gin"
)

// DelTid godoc
// @Summary      删除标签
// @Tags		 标签管理(云平台)
// @Param        id    path	int  true 	"id"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/tid/{id} [delete]
func DelTid(c *gin.Context) {
	var (
		err error
		uri uriId
	)

	if err = c.BindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if db.DeleteOnsInfo(uri.Id); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success"})
}
