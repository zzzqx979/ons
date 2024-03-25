package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/dataPanel/core"
	"ons/util/errors"
)

// QueryCodesByEpc godoc
// @Summary      epc码查询厂商代码和产品代码
// @Tags		 单一查询
// @Accept       json
// @Param		 epc	path	string	true	"epc码"
// @Param		 trace  query	bool  false   "链路追踪参数"
// @Param		 refresh  query	bool  false   "缓存刷新参数"
// @Success      200 {object} response.Response
// @Success		 200 {object} response.TraceResponse
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/epc_codes/{epc} [get]
func QueryCodesByEpc(c *gin.Context) {
	var (
		err   error
		uri   uriEpc
		codes Codes
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

	codes.FactoryCode, codes.ProductCode, trace.Ons, err = core.OM.OnsCodesByEpc(query.Refresh, c.Request.RequestURI, uri.Epc)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if query.Trace {
		c.JSON(http.StatusOK, response.TraceResponse{Code: 200, Message: "success", Data: codes, Trace: trace})
	} else {
		c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success", Data: codes})
	}
}
