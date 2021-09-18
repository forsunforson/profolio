package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/forsunforson/profolio/config"
)

type Stock interface {
	GetName() string
	GetLatestPrice() string
	GetCode() string
	GetDate() string
	GetMarket() string
}

const (
	JuheStockURL = `http://web.juhe.cn:8080/finance/stock/`

	JuheHushen   = "hs"
	JuheHongkong = "hk"
)

type JuheStock struct {
	data *JuheData

	ID        int
	Mtime     time.Time
	StockCode string
	StockName string
	Market    string

	CDate string
	Price string
}

type JuheData struct {
	Gid        string `json:"gid"`
	EName      string `json:"ename"`
	Name       string `json:"name"`
	OpenPri    string `json:"openpri"`
	FormPri    string `json:"formpri"`
	MaxPri     string `json:"maxpri"`
	MinPri     string `json:"minpri"`
	LastestPri string `json:"lastestpri"`
	UpPic      string `json:"uppic"`
	Limit      string `json:"limit"`
	InPic      string `json:"inpic"`
	OutPic     string `json:"outpic"`
	TraAmount  string `json:"traAmount"`
	TraNumber  string `json:"traNumber"`
	PriEarn    string `json:"priearn"`
	Max52      string `json:"max52"`
	Min52      string `json:"min52"`
	Date       string `json:"date"`
	Time       string `json:"time"`
}

type JuheResult struct {
	Data JuheData `json:"data"`
}

type JuheRsp struct {
	ResultCode string       `json:"resultcode"`
	Result     []JuheResult `json:"result"`
}

func NewStock(code string, market string) (Stock, error) {
	params, err := getParams(code, market)
	if err != nil {
		fmt.Printf("unsupported market: %s\n", market)
		return nil, err
	}
	url := fmt.Sprintf("%s%s?%s&key=%s", JuheStockURL, market, params, config.GetGlobalConfig().DateSource[0].AppKey)
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get fail: %v\n", err)
		return nil, err
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Printf("read body fail: %v\n", body)
		return nil, err
	}
	juheRsp := JuheRsp{}
	err = json.Unmarshal(body, &juheRsp)
	if err != nil {
		fmt.Printf("unmarshal body fail: %s\n", err)
		return nil, err
	}
	if juheRsp.ResultCode != "200" {
		fmt.Printf("wrong result code: %s\n", juheRsp.ResultCode)
		return nil, err
	}
	stock := JuheStock{
		data: &juheRsp.Result[0].Data,
	}
	stock.StockCode = stock.data.Gid
	stock.Price = stock.data.LastestPri
	stock.StockName = stock.data.Name
	stock.Market = market
	now := time.Now()
	stock.CDate = fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
	return &stock, nil
}

func getParams(code, market string) (string, error) {
	switch market {
	case JuheHushen:
		return fmt.Sprintf("gid=%s", code), nil
	case JuheHongkong:
		return fmt.Sprintf("num=%s", code), nil
	default:
		return "", errors.New("unsupported")
	}
}

func (s *JuheStock) GetName() string {
	return s.StockName
}

func (s *JuheStock) GetCode() string {
	return s.StockCode
}

func (s *JuheStock) GetLatestPrice() string {
	return s.Price
}

func (s *JuheStock) GetDate() string {
	return s.CDate
}

func (s *JuheStock) GetMarket() string {
	return s.Market
}
