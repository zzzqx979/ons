package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/dataPanel/core"
	"ons/db"
	"ons/util/errors"
	"strconv"
)

// QueryModelByTid godoc
// @Summary      tid查询物模型
// @Tags		 复合查询
// @Accept       json
// @Param		 tid	path	string	true	"tid"
// @Param		 trace  query	bool  false   "链路追踪参数"
// @Param		 refresh  query	bool  false   "缓存刷新参数"
// @Success      200 {object} response.Response
// @Success		 200 {object} response.TraceResponse
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/tid_model/{tid} [get]
func QueryModelByTid(c *gin.Context) {
	var (
		err      error
		uri      uriTid
		tsl      db.TSL
		codesStr string
		trace    response.Trace
		codes    Codes
		query    queryParameter
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
	codesStr = c.GetHeader("codes")
	if codesStr == "" {
		codes.FactoryCode, codes.ProductCode, trace.Ons, err = core.OM.OnsCodesByTid(query.Refresh,
			fmt.Sprintf("/ons/tid_codes/%s?trace=%s&refresh=%s", uri.Tid, strconv.FormatBool(query.Trace), strconv.FormatBool(query.Refresh)), uri.Tid)
		if err != nil {
			logrus.Errorf("QueryModelByTid() ons query errors:%+v.\n", err)
			c.JSON(http.StatusInternalServerError,
				response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
			return
		}
		codesStr = fmt.Sprintf("%s/%s", codes.FactoryCode, codes.ProductCode)
	}
	tsl, trace.Ms, err = core.OM.MsModelByCodes(query.Refresh, c.Request.RequestURI, codesStr)
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
