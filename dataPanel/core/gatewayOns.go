package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"ons/controlPanel/http/response"
	"ons/db"
	"sync"
	"time"
)

type GatewayTidOnsInfo struct {
	FactoryCode    string
	ProductCode    string
	Epc            string
	EpcExpiredAt   time.Time
	UriExpiredAt   time.Time
	CodesExpiredAt time.Time
}

type GatewayEpcOnsInfo struct {
	Tid            string
	FactoryCode    string
	ProductCode    string
	TidExpiredAt   time.Time
	UriExpiredAt   time.Time
	CodesExpiredAt time.Time
}

type GatewayCodesOnsInfo struct {
	Model          db.TSL
	ModelExpiredAt time.Time
}

type Codes struct {
	FactoryCode string `json:"factory_code"`
	ProductCode string `json:"product_code"`
}

type gatewayOM struct {
	onsTidMap   sync.Map
	onsEpcMap   sync.Map
	onsCodesMap sync.Map
}

func (g *gatewayOM) OnsEpcByTid(refresh bool, uri, tid string) (epc string, onsTrace []response.TraceItem, err error) {
	var info GatewayTidOnsInfo
	value, ok := g.onsTidMap.Load(tid)
	if value != nil {
		info, ok = value.(GatewayTidOnsInfo)
		if !ok {
			logrus.Errorf("GatewayOnsByCodes() assert failed. data type is %T.\n", value)
			return
		}
		if refresh {
			info.EpcExpiredAt = time.Now()
		}
	}
	if !ok || value == nil || time.Now().Add(time.Second).After(info.EpcExpiredAt) {
		value, onsTrace, err = QueryOnsUpstream(uri)
		if err != nil {
			logrus.Errorf("GatewayOnsByTid() OnsQueryEpcByTid query ons upstream errors:%+v.\n", err)
			return
		}
		info.Epc, ok = value.(string)
		if !ok {
			logrus.Errorf("GatewayOnsByTid() assert failed. data type is %T.\n", value)
			return
		}
		// 更新过期时间
		info.EpcExpiredAt = time.Now().Add(time.Minute * 30)
		g.onsTidMap.Store(tid, info)
	}
	epc = info.Epc
	return
}

func (g *gatewayOM) OnsCodesByTid(refresh bool, uri, tid string) (factoryCode, productCode string, onsTrace []response.TraceItem, err error) {
	var info GatewayTidOnsInfo
	value, ok := g.onsTidMap.Load(tid)
	if value != nil {
		info, ok = value.(GatewayTidOnsInfo)
		if !ok {
			logrus.Errorf("GatewayOnsByCodes() assert failed. data type is %T.\n", info)
			return
		}
		if refresh {
			info.CodesExpiredAt = time.Now()
		}
	}
	if !ok || value == nil || time.Now().Add(time.Second).After(info.CodesExpiredAt) {
		var (
			codesByte []byte
			codes     Codes
		)
		value, onsTrace, err = QueryOnsUpstream(uri)
		if err != nil {
			logrus.Errorf("GatewayOnsByTid() OnsQueryCodesByTid query ons upstream errors:%+v.\n", err)
			return
		}
		codesByte, err = json.Marshal(value)
		if err != nil {
			logrus.Errorf("GatewayOnsByTid() OnsQueryCodesByTid marshal codes errors:%+v.\n", err)
			return
		}
		err = json.Unmarshal(codesByte, &codes)
		if err != nil {
			logrus.Errorf("GatewayOnsByTid() OnsQueryCodesByTid unmarshal codes errors:%+v.\n", err)
			return
		}
		info.FactoryCode = codes.FactoryCode
		info.ProductCode = codes.ProductCode
		// 更新过期时间
		info.CodesExpiredAt = time.Now().Add(time.Minute * 30)
		g.onsTidMap.Store(tid, info)
	}
	factoryCode = info.FactoryCode
	productCode = info.ProductCode
	return
}

func (g *gatewayOM) OnsTidByEpc(refresh bool, uri, epc string) (tid string, onsTrace []response.TraceItem, err error) {
	var info GatewayEpcOnsInfo
	value, ok := g.onsEpcMap.Load(epc)
	if value != nil {
		info, ok = value.(GatewayEpcOnsInfo)
		if !ok {
			logrus.Errorf("OnsTidByEpc() assert failed. data type is %T.\n", info)
			return
		}
		if refresh {
			info.TidExpiredAt = time.Now()
		}
	}
	if !ok || value == nil || time.Now().Add(time.Second).After(info.TidExpiredAt) {
		value, onsTrace, err = QueryOnsUpstream(uri)
		if err != nil {
			logrus.Errorf("OnsTidByEpc() query ons upstream errors:%+v.\n", err)
			return
		}
		info.Tid, ok = value.(string)
		if !ok {
			logrus.Errorf("OnsTidByEpc() assert failed. data type is %T.\n", info)
			return
		}
		// 更新过期时间
		info.TidExpiredAt = time.Now().Add(time.Minute * 30)
		g.onsEpcMap.Store(epc, info)
	}
	tid = info.Tid
	return
}

func (g *gatewayOM) OnsCodesByEpc(refresh bool, uri, epc string) (factoryCode, productCode string, onsTrace []response.TraceItem, err error) {
	var info GatewayEpcOnsInfo
	value, ok := g.onsEpcMap.Load(epc)
	if value != nil {
		info, ok = value.(GatewayEpcOnsInfo)
		if !ok {
			logrus.Errorf("OnsCodesByEpc() assert failed. data type is %T.\n", info)
			return
		}
		if refresh {
			info.CodesExpiredAt = time.Now()
		}
	}
	if !ok || value == nil || time.Now().Add(time.Second).After(info.CodesExpiredAt) {
		var (
			codesByte []byte
			codes     Codes
		)
		value, onsTrace, err = QueryOnsUpstream(uri)
		if err != nil {
			logrus.Errorf("OnsCodesByEpc() query ons upstream errors:%+v.\n", err)
			return
		}
		codesByte, err = json.Marshal(value)
		if err != nil {
			logrus.Errorf("OnsCodesByEpc() marshal codes errors:%+v.\n", err)
			return
		}
		err = json.Unmarshal(codesByte, &codes)
		if err != nil {
			logrus.Errorf("OnsCodesByEpc() unmarshal codes errors:%+v.\n", err)
			return
		}
		info.FactoryCode = codes.FactoryCode
		info.ProductCode = codes.ProductCode
		// 更新过期时间
		info.CodesExpiredAt = time.Now().Add(time.Minute * 30)
		g.onsEpcMap.Store(epc, info)
	}
	factoryCode = info.FactoryCode
	productCode = info.ProductCode
	return
}

func QueryOnsUpstream(uri string) (data interface{}, onsTrace []response.TraceItem, err error) {
	var (
		url          string
		req          *http.Request
		resp         *http.Response
		dataByte     []byte
		onsUpstreams []string
	)
	onsUpstreams, err = db.GetOnsUpstream()
	if err != nil {
		logrus.Errorf("QueryOnsUpstream() get ons upstreams error:%+v.\n", err)
		return
	}
	for _, upstream := range onsUpstreams {
		if upstream == "" {
			continue
		}
		url = fmt.Sprintf("%s%s", upstream, uri)
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			logrus.Errorf("QueryOnsUpstream() new request error:%+v.\n", err)
			return
		}
		cli := http.Client{}
		st := time.Now()
		// http
		resp, err = cli.Do(req)
		if err != nil {
			logrus.Errorf("QueryOnsUpstream() cli do error:%+v.\n", err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			continue
		}
		et := time.Now()
		dataByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorf("QueryOnsUpstream() read resp body error:%+v.\n", err)
			return
		}
		// json marshal
		v := viper.New()
		v.SetConfigType("json")
		err = v.ReadConfig(bytes.NewBuffer(dataByte))
		if err != nil {
			logrus.Errorf("QueryOnsUpstream() unmarshal GatewayEpcOnsInfo data error:%+v.\n", err)
			return
		}
		data = v.Get("data")
		traceMap := v.GetStringMap("trace")
		dataByte, err = json.Marshal(traceMap["ons"])
		if err != nil {
			logrus.Errorf("QueryOnsUpstream() marshal trace error:%+v.\n", err)
			return
		}
		err = json.Unmarshal(dataByte, &onsTrace)
		if err != nil {
			logrus.Errorf("QueryOnsUpstream() unmarshal trace error:%+v.\n", err)
			return
		}
		onsTrace = append([]response.TraceItem{
			{Upstream: upstream, Time: fmt.Sprintf("%dms", et.UnixMilli()-st.UnixMilli())}}, onsTrace...)
	}
	return
}
