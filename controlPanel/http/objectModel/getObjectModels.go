package objectModel

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

type GetObjectModelsResp struct {
	List  []db.ObjectModel
	Count int64
}

// GetObjectModels godoc
// @Summary      物模型列表
// @Tags		 物模型管理(云平台)
// @Accept       json
// @Param        page    query	int  true 	"页码"
// @Param        page_size    query	int  true 	"页大小"
// @Param        name    query	string  false 	"物模型名称"
// @Param		 product_code	query	string	false	"产品代码"
// @Param		 factory_code	query	string	false	"厂商代码"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/object_model [get]
func GetObjectModels(c *gin.Context) {
	var (
		err   error
		query db.GetObjectModelsQuery
		resp  GetObjectModelsResp
	)

	if err = c.BindQuery(&query); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	resp.Count, resp.List, err = db.GetObjectModels(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success", Data: resp})
}
