package model

type Market string

const (
	Shangzheng Market = "sh"
	Shenzheng  Market = "sz"
	Hongkong   Market = "hk"
)

type Stock struct {
	ID        int
	StockCode string
	StockName string
	Market    string

	LatestPrice string
	OpenPrice   string
	MaxPrice    string
	MinPrice    string
}
