package core

import (
	"ons/controlPanel/http/response"
	"ons/db"
)

var (
	OM OMInterface
)

type OMInterface interface {
	OnsEpcByTid(refresh bool, uri, tid string) (epc string, onsTrace []response.TraceItem, err error)
	OnsCodesByTid(refresh bool, uri, tid string) (factoryCode, productCode string, onsTrace []response.TraceItem, err error)
	OnsTidByEpc(refresh bool, uri, epc string) (tid string, onsTrace []response.TraceItem, err error)
	OnsCodesByEpc(refresh bool, uri, epc string) (factoryCode, productCode string, onsTrace []response.TraceItem, err error)
	MsModelByCodes(refresh bool, uri, codes string) (model db.TSL, msTrace []response.TraceItem, err error)
}

func InitAIotOM() {
	OM = &aIotOM{}
}

func InitGatewayOM() {
	OM = &gatewayOM{}
}
