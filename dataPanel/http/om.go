package http

import (
	"github.com/gin-gonic/gin"
)

type queryParameter struct {
	Trace   bool `form:"trace"`
	Refresh bool `form:"refresh"`
}

type uriTid struct {
	Tid string `uri:"tid"`
}

type uriEpc struct {
	Epc string `uri:"epc"`
}

type UriCodes struct {
	FactoryCode string `uri:"factory_code"`
	ProductCode string `uri:"product_code"`
}

type Codes struct {
	FactoryCode string `json:"factory_code"`
	ProductCode string `json:"product_code"`
}

func RouterOM(root *gin.RouterGroup) {
	group := root.Group("")
	{
		group.GET("/type", ServerType)

		group.GET("tid_epc/:tid", QueryEpcByTid)
		group.GET("tid_codes/:tid", QueryCodesByTid)
		group.GET("tid_uri/:tid", QueryUrlByTid)
		group.GET("tid_model/:tid", QueryModelByTid)

		group.GET("epc_tid/:epc", QueryTidByEpc)
		group.GET("epc_codes/:epc", QueryCodesByEpc)
		group.GET("epc_uri/:epc", QueryUrlByEpc)
		group.GET("epc_model/:epc", QueryModelByEpc)

		group.GET("codes_uri/:factory_code/:product_code", QueryUrlByCodes)
		group.GET("codes_model/:factory_code/:product_code", QueryModelByCodes)
	}
}
