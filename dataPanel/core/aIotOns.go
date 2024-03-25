package core

import (
	"github.com/sirupsen/logrus"
	"ons/controlPanel/http/response"
	"ons/db"
)

var expireTime = 60

type aIotOM struct {
}

func (a *aIotOM) OnsEpcByTid(refresh bool, uri, tid string) (epc string, onsTrace []response.TraceItem, err error) {
	var onsInfo db.OnsInfo
	onsInfo, err = a.OnsByTid(refresh, tid)
	epc = onsInfo.Epc
	return
}

func (a *aIotOM) OnsCodesByTid(refresh bool, uri, tid string) (factoryCode, productCode string, onsTrace []response.TraceItem, err error) {
	var onsInfo db.OnsInfo
	onsInfo, err = a.OnsByTid(refresh, tid)
	factoryCode = onsInfo.FactoryCode
	productCode = onsInfo.ProductCode
	return
}

func (a *aIotOM) OnsByTid(refresh bool, tid string) (onsInfo db.OnsInfo, err error) {
	var onsInfoMap map[string]string
	if !refresh {
		//先查缓存
		cmd := db.GetRedisCli().HGetAll(tid)
		onsInfoMap = cmd.Val()
	}
	//缓存里查不到就去数据库查
	if len(onsInfoMap) == 0 || refresh {
		var has bool
		//TID查询
		onsInfo, has, err = db.GetOnsInfoByTid(tid)
		if !has {
			//缓存空对象
			err = db.AddTidRecordScript.Eval(db.GetRedisCli(),
				[]string{tid, db.RedisFieldEpc, db.RedisFieldFactoryCode, db.RedisFieldProductCode},
				"", "", "", expireTime).Err()
			return
		}
		if err != nil {
			logrus.Errorf("OnsByTid() query ons info err:%+v", err)
			return
		}
		//写入缓存
		err = db.AddTidRecordScript.Eval(db.GetRedisCli(),
			[]string{tid, db.RedisFieldEpc, db.RedisFieldFactoryCode, db.RedisFieldProductCode},
			onsInfo.Epc, onsInfo.FactoryCode, onsInfo.ProductCode, expireTime).Err()
		return
	}
	onsInfo = db.OnsInfo{
		ProductCode: onsInfoMap[db.RedisFieldProductCode],
		FactoryCode: onsInfoMap[db.RedisFieldFactoryCode],
		Epc:         onsInfoMap[db.RedisFieldEpc],
	}
	return
}

func (a *aIotOM) OnsTidByEpc(refresh bool, uri, epc string) (tid string, onsTrace []response.TraceItem, err error) {
	var onsInfo db.OnsInfo
	onsInfo, err = a.OnsByEpc(refresh, epc)
	tid = onsInfo.Tid
	return
}

func (a *aIotOM) OnsCodesByEpc(refresh bool, uri, epc string) (factoryCode, productCode string, onsTrace []response.TraceItem, err error) {
	var onsInfo db.OnsInfo
	onsInfo, err = a.OnsByEpc(refresh, epc)
	factoryCode = onsInfo.FactoryCode
	productCode = onsInfo.ProductCode
	return
}

func (a *aIotOM) OnsByEpc(refresh bool, epc string) (onsInfo db.OnsInfo, err error) {
	var onsInfoMap map[string]string
	if !refresh { //先查缓存
		cmd := db.GetRedisCli().HGetAll(epc)
		onsInfoMap = cmd.Val()
	}
	//缓存里查不到就去数据库查
	if len(onsInfoMap) == 0 || refresh {
		var has bool
		//Epc查询
		onsInfo, has, err = db.GetOnsInfoByEpc(epc)
		if !has {
			//缓存空对象
			err = db.AddEpcRecordScript.Eval(db.GetRedisCli(),
				[]string{epc, db.RedisFieldTid, db.RedisFieldFactoryCode, db.RedisFieldProductCode},
				"", "", "", expireTime).Err()
			return
		}
		if err != nil {
			logrus.Errorf("OnsByEpc() query ons info err:%+v", err)
			return
		}
		//写入缓存
		err = db.AddEpcRecordScript.Eval(db.GetRedisCli(),
			[]string{epc, db.RedisFieldTid, db.RedisFieldFactoryCode, db.RedisFieldProductCode},
			onsInfo.Tid, onsInfo.FactoryCode, onsInfo.ProductCode, expireTime).Err()
		return
	}
	onsInfo = db.OnsInfo{
		ProductCode: onsInfoMap[db.RedisFieldProductCode],
		FactoryCode: onsInfoMap[db.RedisFieldFactoryCode],
		Tid:         onsInfoMap[db.RedisFieldTid],
	}
	return
}
