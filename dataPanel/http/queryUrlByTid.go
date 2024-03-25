package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/util/errors"
)

// QueryUrlByTid godoc
// @Summary      tid查询物模型地址
// @Tags		 单一查询
// @Accept       json
// @Param		 tid	path	string	true	"tid"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/tid_uri/{tid} [get]
func QueryUrlByTid(c *gin.Context) {
	var (
		err error
		uri uriTid
	)

	if err = c.BindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success", Data: fmt.Sprintf("/ons/tid_model/%s", uri.Tid)})
}
