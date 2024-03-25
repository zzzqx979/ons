package mqtt

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"ons/dataPanel/core"
)

const (
	EpcByTid = iota + 1
	CodesByTid
	UriByTid
	ModelByTid
	TidByEpc
	CodesByEpc
	UriByEpc
	ModelByEpc
	UriByCodes
	ModelByCodes
)

type MqttRequest struct {
	Type        int    `json:"type"`
	Tid         string `json:"tid"`
	Epc         string `json:"epc"`
	FactoryCode string `json:"factory_code"`
	ProductCode string `json:"product_code"`
}

type MqttResponse struct {
	Tid       string      `json:"TID"`
	Epc       string      `json:"epc"`
	Uri       string      `json:"uri"`
	FactoryId string      `json:"factoryId"`
	ProductId string      `json:"productId"`
	Model     interface{} `json:"model"`
}

var omCallback mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	payload := MqttRequest{}
	err := json.Unmarshal(msg.Payload(), &payload)
	if err != nil {
		return
	}
	resp := MqttResponse{}
	switch payload.Type {
	case EpcByTid:
		resp.Epc, _, err = core.OM.OnsEpcByTid(false, "", payload.Tid)
	case CodesByTid:
		resp.FactoryId, resp.ProductId, _, err = core.OM.OnsCodesByTid(false, "", payload.Tid)
	case UriByTid:
		resp.Uri = fmt.Sprintf("/ons/tid_model/%s", payload.Tid)
	case ModelByTid:
		var (
			factoryCode string
			productCode string
		)
		factoryCode, productCode, _, err = core.OM.OnsCodesByTid(false, "", payload.Tid)
		if err != nil {
			break
		}
		resp.Model, _, err = core.OM.MsModelByCodes(false, "",
			fmt.Sprintf("%s/%s", factoryCode, productCode))
	case TidByEpc:
		resp.Tid, _, err = core.OM.OnsTidByEpc(false, "", payload.Epc)
	case CodesByEpc:
		resp.FactoryId, resp.ProductId, _, err = core.OM.OnsCodesByEpc(false, "", payload.Epc)
	case UriByEpc:
		resp.Uri = fmt.Sprintf("/ons/epc_model/%s", payload.Epc)
	case ModelByEpc:
		var (
			factoryCode string
			productCode string
		)
		factoryCode, productCode, _, err = core.OM.OnsCodesByEpc(false, "", payload.Epc)
		if err != nil {
			break
		}
		resp.Model, _, err = core.OM.MsModelByCodes(false, "",
			fmt.Sprintf("%s/%s", factoryCode, productCode))
	case UriByCodes:
		resp.Uri = fmt.Sprintf("/ons/codes_model/%s/%s", payload.FactoryCode, payload.ProductCode)
	case ModelByCodes:
		resp.Model, _, err = core.OM.MsModelByCodes(false, "",
			fmt.Sprintf("%s/%s", payload.FactoryCode, payload.ProductCode))
	}
	respByte, err := json.Marshal(resp)
	if err != nil {
		logrus.Errorf("publish response marshal err:%+v.\n", err)
		return
	}
	if token := MqttClient.Publish(responseTopic, 0, false, string(respByte)); token.Error() != nil {
		logrus.Errorf("publish msg to topic:%s err:%+v.\n", responseTopic, err)
	}
}
