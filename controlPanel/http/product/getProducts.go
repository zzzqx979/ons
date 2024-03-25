package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

type GetProductsReq struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Name     string `form:"name"`
	Code     string `form:"code"`
}

type GetProductsResp struct {
	List  []db.Product
	Count int64
}

// GetProducts godoc
// @Summary      产品列表
// @Tags		 产品管理(云平台)
// @Accept       json
// @Param        page    query	int  true 	"页码"
// @Param        page_size    query	int  true 	"页大小"
// @Param        name    query	string  false 	"产品名称"
// @Param        code    query	string  false 	"产品代码"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/product [get]
func GetProducts(c *gin.Context) {
	var (
		err  error
		req  GetProductsReq
		resp GetProductsResp
	)

	if err = c.BindQuery(&req); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	resp.Count, resp.List, err = db.GetProducts(req.Page, req.PageSize, req.Name, req.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success", Data: resp})
}
