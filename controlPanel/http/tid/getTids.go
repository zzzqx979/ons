package tid

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

type GetOnsInfosResp struct {
	List  []db.OnsInfo
	Count int64
}

// GetTids godoc
// @Summary      标签列表
// @Tags		 标签管理(云平台)
// @Accept       json
// @Param        page    query	int  true 	"页码"
// @Param        page_size    query	int  true 	"页大小"
// @Param        tid    query	string  false 	"标签"
// @Param		 product_code	query	string	false	"产品代码"
// @Param		 factory_code	query	string	false	"厂商代码"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/tid [get]
func GetTids(c *gin.Context) {
	var (
		err   error
		query db.GetOnsInfosQuery
		resp  GetOnsInfosResp
	)

	if err = c.BindQuery(&query); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	resp.Count, resp.List, err = db.GetOnsInfos(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success", Data: resp})
}
