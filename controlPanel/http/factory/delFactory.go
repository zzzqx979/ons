package factory

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
	"strconv"
)

// DelFactory godoc
// @Summary      删除厂商
// @Tags		 厂商管理(云平台)
// @Accept       json
// @Param        id    path	int  true 	"厂商id"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/factory/{id} [delete]
func DelFactory(c *gin.Context) {
	var (
		err error
		id  int
	)

	idStr := c.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if err = db.DelFactory(id); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 0, Message: "success"})
}
