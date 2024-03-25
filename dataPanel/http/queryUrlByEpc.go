package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/util/errors"
)

// QueryUrlByEpc godoc
// @Summary      epc码查询物模型地址
// @Tags		 单一查询
// @Accept       json
// @Param		 epc	path	string	true	"epc码"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/epc_uri/{epc} [get]
func QueryUrlByEpc(c *gin.Context) {
	var (
		err error
		uri uriEpc
	)

	if err = c.BindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success", Data: fmt.Sprintf("/ons/epc_model/%s", uri.Epc)})
}
