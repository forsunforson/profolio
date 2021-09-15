package model

const (
	JuheExchangeURL = "http://op.juhe.cn/onebox/exchange/query"
)

type ExchangeRate interface {
	USD2RMB(float64) int
	HKD2RMB(float64) int
}
