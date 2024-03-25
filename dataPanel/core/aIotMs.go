package core

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"ons/controlPanel/http/response"
	"ons/db"
	"strings"
)

func (a *aIotOM) MsModelByCodes(refresh bool, uri, codes string) (tsl db.TSL, msTrace []response.TraceItem, err error) {
	var tslStr string
	if !refresh {
		//先查缓存
		cmd := db.GetRedisCli().Get(codes)
		tslStr = cmd.Val()
	}
	if tslStr == "" || refresh {
		var has bool
		tsl, has, err = db.GetObjectModelByCodes(strings.Split(codes, "/")[0], strings.Split(codes, "/")[1])
		if !has {
			err = db.AddCodesRecordScript.Eval(db.GetRedisCli(),
				[]string{codes}, "", expireTime).Err()
			return
		}
		if err != nil {
			logrus.Errorf("MsModelByCodes() get get object modelStr errors:%+v.\n", err)
			return
		}
		tsl = db.TSL{
			Version: "v1.0",
			Profile: db.Profile{
				ProductCode: tsl.ProductCode,
				FactoryCode: tsl.FactoryCode,
			},
			Properties: tsl.Properties,
			Events:     tsl.Events,
			Services:   tsl.Services,
		}
		var tslByte []byte
		tslByte, err = json.Marshal(tsl)
		if err != nil {
			logrus.Errorf("MsModelByCodes() marshal model err:%+v.\n", err)
			return
		}
		err = db.AddCodesRecordScript.Eval(db.GetRedisCli(),
			[]string{codes}, string(tslByte), expireTime).Err()
		return
	}
	err = json.Unmarshal([]byte(tslStr), &tsl)
	if err != nil {
		logrus.Errorf("MsModelByCodes() unmarshal modelStr error:%+v.\n", err)
	}
	return
}
