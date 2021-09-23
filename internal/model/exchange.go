package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/forsunforson/profolio/config"
	"github.com/golang/glog"
)

const (
	JuheExchangeURL = "http://op.juhe.cn/onebox/exchange/query"
)

type ExchangeRate interface {
	USD2RMB(float64) int
	HKD2RMB(float64) int

	Update()
}

type JuheExchange struct {
	usd2RMBRate float64
	rmb2USDRate float64
	hdk2RMBRate float64
	rmb2HKDRate float64
	update      time.Time
	result      *JuheExResult
	l           sync.Mutex
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

func NewExchangeRate() ExchangeRate {
	ret := JuheExchange{
		update: time.Now().Add(-24 * time.Hour),
	}
	ret.Update()
	parseExRate(&ret)
	return &ret
}

func parseExRate(exchange *JuheExchange) {
	for _, v := range exchange.result.List {
		if len(v) < 6 {
			continue
		}
		midBid := v[5]
		rate, _ := strconv.ParseFloat(midBid, 64)
		switch v[0] {
		case "美元":
			exchange.usd2RMBRate = rate
			exchange.rmb2USDRate = 1.0 / rate
		case "港币":
			exchange.hdk2RMBRate = rate
			exchange.rmb2HKDRate = 1.0 / rate
		default:
			continue
		}
	}
}

func (e *JuheExchange) USD2RMB(usd float64) int {
	rmb := int(usd * e.usd2RMBRate)
	return rmb
}

func (e *JuheExchange) HKD2RMB(hkd float64) int {
	rmb := int(hkd * e.hdk2RMBRate)
	return rmb
}

func (e *JuheExchange) Update() {
	now := time.Now()
	if e.update.Day() == now.Day() {
		return
	}
	url := fmt.Sprintf(JuheExchangeURL+"?key=%s", config.GetGlobalConfig().DateSource[0].ExchangeKey)
	rsp, err := http.Get(url)
	if err != nil {
		glog.Errorf("http get fail: %v", err)
		return
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		glog.Errorf("read body fail: %v", body)
		return
	}
	juheRsp := JuheExRsp{}
	err = json.Unmarshal(body, &juheRsp)
	if err != nil {
		glog.Errorf("unmarshal body fail: %s", err)
		return
	}
	if juheRsp.ErrorCode != 0 {
		glog.Errorf("error code is %d: %s", juheRsp.ErrorCode, juheRsp.Reason)
		return
	}
	e.l.Lock()
	defer e.l.Unlock()
	e.result = &juheRsp.Result
	e.update = now
}
