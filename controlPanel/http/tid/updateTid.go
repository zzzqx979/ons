package tid

import (
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"

	"github.com/gin-gonic/gin"
)

// UpdateTid godoc
// @Summary      更新标签
// @Tags		 标签管理(云平台)
// @Accept       json
// @Param        id    body	int  true 	"标签id"
// @Param        tid    body	string  true 	"标签"
// @Param        epc    body	string  true 	"epc码"
// @Param        url    body	string  true 	"物模型url"
// @Param		 product_code	body	string	true	"产品代码"
// @Param		 factory_code	body	string	true	"厂商代码"
// @Param		 note	body	string	false	"备注"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/tid [put]
func UpdateTid(c *gin.Context) {
	var (
		err error
		req db.OnsInfo
	)

	if err = c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if req.FactoryCode == "0" || req.ProductCode == "0" || req.Tid == "0" || req.Epc == "0" {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.ValueZeroErr, Message: errors.New(errors.ValueZeroErr).Error()})
		return
	}

	if err = db.UpdateOnsInfo(req); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success"})
}
