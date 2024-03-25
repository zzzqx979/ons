package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

// UpdateProduct godoc
// @Summary      更新产品
// @Tags		 产品管理(云平台)
// @Accept       json
// @Param		id body int true "产品id"
// @Param        name    body	string  true 	"产品名称"
// @Param		 code	body	string	true	"产品代码"
// @Param		 status	body	int	true	"状态 1是启用 0是禁用"
// @Param		 note	body	string	false	"备注"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/product [put]
func UpdateProduct(c *gin.Context) {
	var (
		err error
		req db.Product
	)

	if err = c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if req.Code == "0" {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.ValueZeroErr, Message: errors.New(errors.ValueZeroErr).Error()})
		return
	}

	if err = db.UpdateProduct(req); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success"})
}
