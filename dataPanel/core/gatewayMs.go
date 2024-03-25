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
	"time"
)

func (g *gatewayOM) MsModelByCodes(refresh bool, uri, codes string) (model db.TSL, msTrace []response.TraceItem, err error) {
	var info GatewayCodesOnsInfo
	value, ok := g.onsCodesMap.Load(codes)
	if value != nil {
		info, ok = value.(GatewayCodesOnsInfo)
		if !ok {
			logrus.Errorf("MsModelByCodes() assert failed,data type is %T\n", value)
			return
		}
		if refresh {
			info.ModelExpiredAt = time.Now()
		}
	}
	if !ok || value == nil || time.Now().Add(time.Second).After(info.ModelExpiredAt) {
		info.Model, msTrace, err = g.QueryMsUpstream(uri, codes)
		if err != nil {
			logrus.Errorf("MsModelByCodes() query ms upstream errors:%+v.\n", err)
			return
		}
		info.ModelExpiredAt = time.Now().Add(time.Minute * 30)
		g.onsCodesMap.Store(codes, info)
	}
	model = info.Model
	return
}

func (g *gatewayOM) QueryMsUpstream(uri, codes string) (tsl db.TSL, msTrace []response.TraceItem, err error) {
	var (
		url         string
		req         *http.Request
		resp        *http.Response
		dataByte    []byte
		msUpstreams []string
	)
	msUpstreams, err = db.GetMsUpstream()
	if err != nil {
		logrus.Errorf("QueryMsUpstream() get ons upstreams error:%+v.\n", err)
		return
	}
	for _, upstream := range msUpstreams {
		if upstream == "" {
			continue
		}
		url = fmt.Sprintf("%s%s", upstream, uri)
		req, err = http.NewRequest("GET", url, nil)
		if err != nil {
			logrus.Errorf("QueryMsUpstream() new request error:%+v.\n", err)
			return
		}
		req.Header.Add("codes", codes)
		cli := http.Client{}
		st := time.Now()
		// http
		resp, err = cli.Do(req)
		if err != nil {
			logrus.Errorf("QueryMsUpstream() cli do error:%+v.\n", err)
			return
		}
		if resp.StatusCode != http.StatusOK {
			continue
		}
		et := time.Now()
		dataByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Errorf("QueryMsUpstream() read resp body error:%+v.\n", err)
			return
		}
		// json marshal
		v := viper.New()
		v.SetConfigType("json")
		err = v.ReadConfig(bytes.NewBuffer(dataByte))
		if err != nil {
			logrus.Errorf("QueryMsUpstream() unmarshal GatewayEpcOnsInfo data error:%+v.\n", err)
			return
		}
		dataByte, err = json.Marshal(v.GetStringMap("data"))
		if err != nil {
			logrus.Errorf("QueryMsUpstream() marshal tsl error:%+v.\n", err)
			return
		}
		err = json.Unmarshal(dataByte, &tsl)
		if err != nil {
			logrus.Errorf("QueryMsUpstream() unmarshal tsl error:%+v.\n", err)
			return
		}
		traceMap := v.GetStringMap("trace")
		dataByte, err = json.Marshal(traceMap["ms"])
		if err != nil {
			logrus.Errorf("QueryMsUpstream() marshal trace error:%+v.\n", err)
			return
		}
		err = json.Unmarshal(dataByte, &msTrace)
		if err != nil {
			logrus.Errorf("QueryMsUpstream() marshal trace error:%+v.\n", err)
			return
		}
		msTrace = append([]response.TraceItem{
			{Upstream: upstream, Time: fmt.Sprintf("%dms", et.UnixMilli()-st.UnixMilli())}}, msTrace...)
	}
	return
}
