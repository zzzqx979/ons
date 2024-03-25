package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/dataPanel/core"
	"ons/db"
	"ons/util/errors"
)

// QueryModelByCodes godoc
// @Summary      厂商代码和产品代码查询物模型
// @Tags		 单一查询
// @Accept       json
// @Param		 factory_code	path	string	true	"厂商代码"
// @Param		 product_code	path	string	true	"产品代码"
// @Param		 trace  query	bool  false   "链路追踪参数"
// @Param		 refresh  query	bool  false   "缓存刷新参数"
// @Success      200 {object} response.Response
// @Success		 200 {object} response.TraceResponse
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/codes_model/{factory_code}/{product_code} [get]
func QueryModelByCodes(c *gin.Context) {
	var (
		err   error
		uri   UriCodes
		tsl   db.TSL
		trace response.Trace
		query queryParameter
	)

	if err = c.BindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if err = c.BindQuery(&query); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	tsl, trace.Ms, err = core.OM.MsModelByCodes(query.Refresh, c.Request.RequestURI, fmt.Sprintf("%s/%s", uri.FactoryCode, uri.ProductCode))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if query.Trace {
		c.JSON(http.StatusOK, response.TraceResponse{Code: 200, Message: "success", Data: tsl, Trace: trace})
	} else {
		c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success", Data: tsl})
	}
}
