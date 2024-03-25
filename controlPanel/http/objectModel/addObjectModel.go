package objectModel

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

type ObjectModelRequest struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Status      int    `json:"status"`
	Note        string `json:"note"`
	ProductCode string `json:"product_code"`
	FactoryCode string `json:"factory_code"`
	Properties  string `json:"properties"`
	Events      string `json:"events"`
	Services    string `json:"services"`
}

// AddObjectModel godoc
// @Summary      添加物模型
// @Tags		 物模型管理(云平台)
// @Accept       json
// @Param        request body ObjectModelRequest true "请求参数"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/object_model [post]
func AddObjectModel(c *gin.Context) {
	var (
		err   error
		req   ObjectModelRequest
		model db.ObjectModel
	)

	if err = c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}

	if req.FactoryCode == "0" || req.ProductCode == "0" {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.ValueZeroErr, Message: errors.New(errors.ValueZeroErr).Error()})
		return
	}
	model = db.ObjectModel{
		Name:        req.Name,
		Status:      req.Status,
		ProductCode: req.ProductCode,
		FactoryCode: req.FactoryCode,
		Properties:  req.Properties,
		Events:      req.Events,
		Services:    req.Services,
		Note:        req.Note,
	}
	if err = db.AddObjectModel(model); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success"})
}
