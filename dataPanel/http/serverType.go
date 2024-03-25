package http

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"ons/controlPanel/http/response"
)

// ServerType godoc
// @Summary      服务类型
// @Tags		 服务类型
// @Accept       json
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/type [get]
func ServerType(c *gin.Context) {
	if viper.IsSet("redis") {
		c.JSON(http.StatusOK, response.TraceResponse{Code: 0, Message: "success", Data: "aiot"})
	} else {
		c.JSON(http.StatusOK, response.TraceResponse{Code: 0, Message: "success", Data: "gateway"})
	}
}
