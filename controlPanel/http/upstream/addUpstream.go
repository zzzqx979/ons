package upstream

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

// AddUpstream godoc
// @Summary      添加上游地址
// @Tags		 upstreams(边缘网关)
// @Accept       json
// @Param        ons1    body	string  true 	"首选上游标签服务器"
// @Param        ons2    body	string  false 	"次选上游标签服务器"
// @Param        ons3    body	string  false 	"备选上游标签服务器"
// @Param        ms1    body	string  true 	"首选上游物模型服务器"
// @Param        ms2    body	string  false 	"次选上游物模型服务器"
// @Param        ms3    body	string  false 	"备选上游物模型服务器"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/upstream [post]
func AddUpstream(c *gin.Context) {
	var (
		upstream db.Upstream
		err      error
	)

	if err = c.BindJSON(&upstream); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if err = db.AddUpstream(upstream); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 0, Message: "success"})
}
