package upstream

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

// GetUpstream godoc
// @Summary      获取上游地址
// @Tags		 upstreams(边缘网关)
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/upstream [get]
func GetUpstream(c *gin.Context) {
	var (
		upstream db.Upstream
		err      error
	)

	if upstream, err = db.GetUpstream(); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 0, Message: "success", Data: upstream})
}
