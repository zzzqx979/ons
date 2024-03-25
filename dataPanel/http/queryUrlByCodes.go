package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/util/errors"
)

// QueryUrlByCodes godoc
// @Summary      厂商代码和产品代码查询物模型地址
// @Tags		 单一查询
// @Accept       json
// @Param		 factory_code	path	string	true	"厂商代码"
// @Param		 product_code	path	string	true	"产品代码"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/codes_uri/{factory_code}/{product_code} [get]
func QueryUrlByCodes(c *gin.Context) {
	var (
		err error
		uri UriCodes
	)

	if err = c.BindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success",
		Data: fmt.Sprintf("/ons/codes_model/%s/%s", uri.FactoryCode, uri.ProductCode)})
}
