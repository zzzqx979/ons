package upstream

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

type DelUpstreamUri struct {
	Id int `uri:"id"`
}

// DelUpstream godoc
// @Summary      删除上游地址
// @Tags		 upstreams(边缘网关)
// @Accept       json
// @Param        id    path	int  true 	"id"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/upstream/{id} [delete]
func DelUpstream(c *gin.Context) {
	var (
		uri DelUpstreamUri
		err error
	)
	if err = c.BindUri(&uri); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if err = db.DeleteUpstream(uri.Id); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 0, Message: "success"})
}
