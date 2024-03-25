package objectModel

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"ons/util/errors"
)

// UpdateObjectModel godoc
// @Summary      更新物模型
// @Tags		 物模型管理(云平台)
// @Accept       json
// @Param        request body ObjectModelRequest true "请求参数"
// @Success      200 {object} response.Response
// @Failure      400 {object} response.Response
// @Failure      404
// @Failure      500 {object} response.Response
// @Router       /ons/object_model [put]
func UpdateObjectModel(c *gin.Context) {
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
		Id:          req.Id,
		Name:        req.Name,
		Status:      req.Status,
		ProductCode: req.ProductCode,
		FactoryCode: req.FactoryCode,
		Properties:  req.Properties,
		Events:      req.Events,
		Services:    req.Services,
		Note:        req.Note,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	model.Id = req.Id
	if err = db.UpdateObjectModel(model); err != nil {
		c.JSON(http.StatusInternalServerError,
			response.Response{Code: errors.InternalErr, Message: errors.New(errors.InternalErr).Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{Code: 200, Message: "success"})
}
