package logic

import (
	"github.com/forsunforson/profolio/internal/model"
)

type RuntimeContext struct {
	Portfolios []*model.Portfolio
	Stocks     map[string]model.Stock
}

var (
	runtimeContext *RuntimeContext
)

func InitContext() {
	ctx := RuntimeContext{}
	stocks := GetAllStocks()
	stockMap := make(map[string]model.Stock)
	for _, stock := range stocks {
		stockMap[stock.GetCode()] = stock
	}
	ctx.Stocks = stockMap
	runtimeContext = &ctx
	go Ticker()
}

func GetRunTimeContext() *RuntimeContext {
	return runtimeContext
}
