package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/forsunforson/profolio/config"
	"github.com/golang/glog"
)

const (
	JuheExchangeURL = "http://op.juhe.cn/onebox/exchange/query"
)

type ExchangeRate interface {
	USD2RMB(float64) int
	HKD2RMB(float64) int
}

type JuheExchange struct {
	usd2RMBRate float64
	rmb2USDRate float64
	hdk2RMBRate float64
	rmb2HKDRate float64
	result      *JuheExResult
}

type JuheExResult struct {
	Update string     `json:"update"`
	List   [][]string `json:"list"`
}

type JuheExRsp struct {
	Reason    string       `json:"reason"`
	Result    JuheExResult `json:"result"`
	ErrorCode int          `json:"error_code"`
}

func NewExchangeRate() (ExchangeRate, error) {
	url := fmt.Sprintf(JuheExchangeURL+"?key=%s", config.GetGlobalConfig().DateSource[0].AppKey)
	rsp, err := http.Get(url)
	if err != nil {
		glog.Errorf("http get fail: %v", err)
		return nil, err
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		glog.Errorf("read body fail: %v", body)
		return nil, err
	}
	juheRsp := JuheExRsp{}
	err = json.Unmarshal(body, &juheRsp)
	if err != nil {
		glog.Errorf("unmarshal body fail: %s", err)
		return nil, err
	}
	if juheRsp.ErrorCode != 0 {
		glog.Errorf("error code is %d: %s", juheRsp.ErrorCode, juheRsp.Reason)
		return nil, err
	}
	ret := JuheExchange{
		result: &juheRsp.Result,
	}
	parseExRate(&ret)
	//TODO解析出汇率进行备用
	return &ret, nil
}

func parseExRate(exchange *JuheExchange) {

}

func (e *JuheExchange) USD2RMB(usd float64) int {
	return 0
}

func (e *JuheExchange) HKD2RMB(hkd float64) int {
	return 0
}
